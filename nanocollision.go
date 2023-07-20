package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/evanj/hacks/trivialstats"
)

func main() {
	numConsecutive := flag.Int("numConsecutive", 5000,
		"number of consecutive time samples to measure at program start")
	minInterestingDiff := flag.Duration("minInterestingDiff", 100*time.Nanosecond,
		"print sample diffs larger than this")
	printFirst := flag.Int("printFirst", 100, "print this many initial diffs")
	longerZeroTest := flag.Int("longerZeroTest", 10000000, "number of samples looking for zero")
	goroutines := flag.Int("goroutines", runtime.GOMAXPROCS(0),
		"number of goroutines to run with the parallel test")
	flag.Parse()

	samples := make([]time.Time, *numConsecutive)
	for i := 0; i < *numConsecutive; i++ {
		samples[i] = time.Now()
	}

	lastTime := samples[0]
	timeDiffDistribution := trivialstats.NewDistribution()
	nanoDiffDistribution := trivialstats.NewDistribution()
	for i, next := range samples[1:] {
		timeDiff := next.Sub(lastTime)
		nanoDiff := next.UnixNano() - lastTime.UnixNano()
		if i < *printFirst || timeDiff <= 0 || timeDiff >= *minInterestingDiff {
			fmt.Printf("%d: %d: timeDiff=%d; nanoDiff=%d\n",
				i+1, next.UnixNano(), timeDiff.Nanoseconds(), nanoDiff)
		}
		lastTime = next
		timeDiffDistribution.Add(timeDiff.Nanoseconds())
		nanoDiffDistribution.Add(nanoDiff)
	}
	fmt.Printf("printed the first %d diffs, and diffs == 0 || >= %s; see --help for flags\n",
		*printFirst, minInterestingDiff)
	fmt.Printf("time diff distribution: %s\n", timeDiffDistribution.Stats())
	fmt.Printf("nano diff distribution: %s\n", nanoDiffDistribution.Stats())

	fmt.Printf("\nrunning longer zeros test ...\n")
	timeDiffZeroCount := 0
	nanoDiffZeroCount := 0
	for i := 0; i < *longerZeroTest; i++ {
		s1 := time.Now()
		s2 := time.Now()
		timeDiff := s2.Sub(s1)
		if timeDiff <= 0 {
			timeDiffZeroCount += 1
			if timeDiff < 0 {
				fmt.Printf("*** negative diff=%d: s1=%s s2=%s", timeDiff, s1, s2)
			}
		}
		nanoDiff := s2.UnixNano() - s1.UnixNano()
		if nanoDiff <= 0 {
			nanoDiffZeroCount += 1
			if nanoDiff < 0 {
				fmt.Printf("*** negative diff=%d: s1=%s s2=%s", nanoDiff, s1, s2)
			}
		}
	}
	fmt.Printf("sampled %d pairs; %d time diff zeros = %f%%; %d nano diff zeros = %f%%\n",
		*longerZeroTest, timeDiffZeroCount, float64(timeDiffZeroCount)*100.0/float64(*longerZeroTest),
		nanoDiffZeroCount, float64(nanoDiffZeroCount)*100.0/float64(*longerZeroTest),
	)

	fmt.Printf("\nstarting parallel test %d goroutines x %d samples ...\n",
		*goroutines, *longerZeroTest)
	results := make(chan *roaring64.Bitmap)
	for i := 0; i < *goroutines; i++ {
		go recordTimestamps(results, *longerZeroTest)
	}

	total := roaring64.New()
	totalCollisions := 0
	for i := 0; i < *goroutines; i++ {
		next := <-results

		numCollisions := next.AndCardinality(total)
		fmt.Printf("%d samples from a thread; %d collisions inside the thread; %d collisions with other threads\n",
			next.GetCardinality(), *longerZeroTest-int(next.GetCardinality()), numCollisions)
		if numCollisions > 0 {
			totalCollisions += int(numCollisions)
		}
		total.Or(next)
	}
	fmt.Printf("%d final samples; %d total collisions = %f%%; possible duplicate collisions? %d\n",
		total.GetCardinality(), totalCollisions,
		float64(totalCollisions)*100.0/float64(*goroutines**longerZeroTest),
		totalCollisions-(*goroutines**longerZeroTest-int(total.GetCardinality())),
	)
}

func recordTimestamps(resultsChan chan<- *roaring64.Bitmap, numSamples int) {
	samples := make([]time.Time, numSamples)
	for i := range samples {
		samples[i] = time.Now()
	}

	results := roaring64.NewBitmap()
	for _, sample := range samples {
		results.Add(uint64(sample.UnixNano()))
	}
	resultsChan <- results
}
