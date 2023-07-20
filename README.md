# NanoCollision: Investigating nanosecond collisions

How often do nanosecond counters collide? Often! See [my blog post for details](https://www.evanjones.ca/nanosecond-collisions.html).

Go's `time.Now()` API records both the "absolute" time as well as a "monotonic clock" for relative time. See [the Go time package documentation for details](https://pkg.go.dev/time).

## Example output

From a Linux system running Ubuntu 23.04, kernel version 6.2.0-25-generic, on real non-virtualized hardware "11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz"

```
1: 1689888428505845688: timeDiff=59; nanoDiff=64
2: 1689888428505845722: timeDiff=34; nanoDiff=34
3: 1689888428505845756: timeDiff=35; nanoDiff=34
4: 1689888428505845789: timeDiff=33; nanoDiff=33
5: 1689888428505845825: timeDiff=35; nanoDiff=36
6: 1689888428505845858: timeDiff=34; nanoDiff=33
7: 1689888428505845892: timeDiff=33; nanoDiff=34
8: 1689888428505845927: timeDiff=36; nanoDiff=35
9: 1689888428505845962: timeDiff=33; nanoDiff=35
10: 1689888428505845994: timeDiff=34; nanoDiff=32
11: 1689888428505846030: timeDiff=34; nanoDiff=36
12: 1689888428505846063: timeDiff=34; nanoDiff=33
13: 1689888428505846098: timeDiff=34; nanoDiff=35
14: 1689888428505846131: timeDiff=34; nanoDiff=33
15: 1689888428505846165: timeDiff=34; nanoDiff=34
16: 1689888428505846199: timeDiff=35; nanoDiff=34
17: 1689888428505846232: timeDiff=31; nanoDiff=33
18: 1689888428505846264: timeDiff=33; nanoDiff=32
19: 1689888428505846297: timeDiff=33; nanoDiff=33
20: 1689888428505846330: timeDiff=33; nanoDiff=33
21: 1689888428505846362: timeDiff=32; nanoDiff=32
22: 1689888428505846394: timeDiff=33; nanoDiff=32
23: 1689888428505846428: timeDiff=31; nanoDiff=34
24: 1689888428505846459: timeDiff=34; nanoDiff=31
25: 1689888428505846493: timeDiff=32; nanoDiff=34
26: 1689888428505846524: timeDiff=33; nanoDiff=31
27: 1689888428505846557: timeDiff=32; nanoDiff=33
28: 1689888428505846589: timeDiff=33; nanoDiff=32
29: 1689888428505846623: timeDiff=32; nanoDiff=34
30: 1689888428505846654: timeDiff=33; nanoDiff=31
31: 1689888428505846687: timeDiff=32; nanoDiff=33
32: 1689888428505846719: timeDiff=33; nanoDiff=32
33: 1689888428505846752: timeDiff=32; nanoDiff=33
34: 1689888428505846785: timeDiff=33; nanoDiff=33
35: 1689888428505846817: timeDiff=31; nanoDiff=32
36: 1689888428505846850: timeDiff=34; nanoDiff=33
37: 1689888428505846882: timeDiff=32; nanoDiff=32
38: 1689888428505846915: timeDiff=33; nanoDiff=33
39: 1689888428505846947: timeDiff=32; nanoDiff=32
40: 1689888428505846980: timeDiff=33; nanoDiff=33
41: 1689888428505847013: timeDiff=33; nanoDiff=33
42: 1689888428505847044: timeDiff=32; nanoDiff=31
43: 1689888428505847078: timeDiff=33; nanoDiff=34
44: 1689888428505847109: timeDiff=31; nanoDiff=31
45: 1689888428505847142: timeDiff=33; nanoDiff=33
46: 1689888428505847175: timeDiff=33; nanoDiff=33
47: 1689888428505847208: timeDiff=32; nanoDiff=33
48: 1689888428505847239: timeDiff=33; nanoDiff=31
49: 1689888428505847272: timeDiff=32; nanoDiff=33
50: 1689888428505847305: timeDiff=32; nanoDiff=33
51: 1689888428505847337: timeDiff=33; nanoDiff=32
52: 1689888428505847370: timeDiff=33; nanoDiff=33
53: 1689888428505847402: timeDiff=32; nanoDiff=32
54: 1689888428505847434: timeDiff=33; nanoDiff=32
55: 1689888428505847468: timeDiff=32; nanoDiff=34
56: 1689888428505847499: timeDiff=32; nanoDiff=31
57: 1689888428505847533: timeDiff=34; nanoDiff=34
58: 1689888428505847565: timeDiff=32; nanoDiff=32
59: 1689888428505847598: timeDiff=32; nanoDiff=33
60: 1689888428505847629: timeDiff=33; nanoDiff=31
61: 1689888428505847663: timeDiff=33; nanoDiff=34
62: 1689888428505847695: timeDiff=32; nanoDiff=32
63: 1689888428505847727: timeDiff=32; nanoDiff=32
64: 1689888428505847759: timeDiff=33; nanoDiff=32
65: 1689888428505847793: timeDiff=33; nanoDiff=34
66: 1689888428505847825: timeDiff=32; nanoDiff=32
67: 1689888428505847857: timeDiff=32; nanoDiff=32
68: 1689888428505847890: timeDiff=34; nanoDiff=33
69: 1689888428505847923: timeDiff=32; nanoDiff=33
70: 1689888428505847955: timeDiff=32; nanoDiff=32
71: 1689888428505847988: timeDiff=32; nanoDiff=33
72: 1689888428505848020: timeDiff=33; nanoDiff=32
73: 1689888428505848053: timeDiff=33; nanoDiff=33
74: 1689888428505848085: timeDiff=32; nanoDiff=32
75: 1689888428505848117: timeDiff=32; nanoDiff=32
76: 1689888428505848150: timeDiff=33; nanoDiff=33
77: 1689888428505848183: timeDiff=33; nanoDiff=33
78: 1689888428505848215: timeDiff=32; nanoDiff=32
79: 1689888428505848247: timeDiff=33; nanoDiff=32
80: 1689888428505848280: timeDiff=32; nanoDiff=33
81: 1689888428505848313: timeDiff=33; nanoDiff=33
82: 1689888428505848345: timeDiff=32; nanoDiff=32
83: 1689888428505848378: timeDiff=33; nanoDiff=33
84: 1689888428505848410: timeDiff=32; nanoDiff=32
85: 1689888428505848443: timeDiff=32; nanoDiff=33
86: 1689888428505848475: timeDiff=33; nanoDiff=32
87: 1689888428505848508: timeDiff=32; nanoDiff=33
88: 1689888428505848542: timeDiff=35; nanoDiff=34
89: 1689888428505848574: timeDiff=32; nanoDiff=32
90: 1689888428505848607: timeDiff=32; nanoDiff=33
91: 1689888428505848639: timeDiff=33; nanoDiff=32
92: 1689888428505848672: timeDiff=33; nanoDiff=33
93: 1689888428505848704: timeDiff=32; nanoDiff=32
94: 1689888428505848736: timeDiff=33; nanoDiff=32
95: 1689888428505848770: timeDiff=32; nanoDiff=34
96: 1689888428505848802: timeDiff=33; nanoDiff=32
97: 1689888428505848835: timeDiff=33; nanoDiff=33
98: 1689888428505848866: timeDiff=32; nanoDiff=31
99: 1689888428505848900: timeDiff=33; nanoDiff=34
100: 1689888428505848931: timeDiff=31; nanoDiff=31
299: 1689888428505860488: timeDiff=5115; nanoDiff=5117
553: 1689888428505867935: timeDiff=315; nanoDiff=30
554: 1689888428505868766: timeDiff=755; nanoDiff=831
555: 1689888428505869284: timeDiff=476; nanoDiff=518
680: 1689888428505874658: timeDiff=1180; nanoDiff=1178
723: 1689888428505887046: timeDiff=11022; nanoDiff=11024
4899: 1689888428506023920: timeDiff=1155; nanoDiff=1155
4904: 1689888428506027957: timeDiff=3906; nanoDiff=3907
printed the first 100 diffs, and diffs == 0 || >= 100ns; see --help for flags
time diff distribution: count=4999 avg=37.0 min=28 p50=32 p90=33 p95=34 max=11022
nano diff distribution: count=4999 avg=37.0 min=28 p50=32 p90=33 p95=34 max=11024

running longer zeros test ...
sampled 10000000 pairs; 0 time diff zeros = 0.000000%; 0 nano diff zeros = 0.000000%

starting parallel test 8 goroutines x 10000000 samples ...
10000000 samples from a thread; 0 collisions inside the thread; 0 collisions with other threads
10000000 samples from a thread; 0 collisions inside the thread; 188465 collisions with other threads
10000000 samples from a thread; 0 collisions inside the thread; 313291 collisions with other threads
10000000 samples from a thread; 0 collisions inside the thread; 497677 collisions with other threads
10000000 samples from a thread; 0 collisions inside the thread; 660498 collisions with other threads
10000000 samples from a thread; 0 collisions inside the thread; 798538 collisions with other threads
10000000 samples from a thread; 0 collisions inside the thread; 932329 collisions with other threads
10000000 samples from a thread; 0 collisions inside the thread; 1099527 collisions with other threads
75509675 final samples; 4490325 total collisions = 5.612906%; possible duplicate collisions? 0

```