#include <stdio.h>
#include <time.h>
#include <inttypes.h>
#include <assert.h>

static const size_t NUM_SAMPLES = 30;

struct clock_description
{
    const char *name;
    clockid_t clock_id;
};

static const struct clock_description CLOCKS[] = {
    {"CLOCK_REALTIME", CLOCK_REALTIME},
    {"CLOCK_MONOTONIC", CLOCK_MONOTONIC},
    // not POSIX but included in both Mac OS X and Linux
    {"CLOCK_MONOTONIC_RAW", CLOCK_MONOTONIC_RAW},
};
static const size_t NUM_CLOCKS = sizeof(CLOCKS) / sizeof(*CLOCKS);

int64_t
timespec_to_nanos(const struct timespec *t)
{
    return t->tv_sec * 1000000000 + t->tv_nsec;
}

int64_t clock_nanos(clockid_t clock_id)
{
    struct timespec out;
    int result = clock_gettime(clock_id, &out);
    assert(result == 0);
    return timespec_to_nanos(&out);
}

int main()
{
    for (int i = 0; i < NUM_CLOCKS; i++)
    {
        struct timespec out;
        int result = clock_getres(CLOCKS[i].clock_id, &out);
        assert(result == 0);
        printf("clock_getres(%s, ...)=%" PRId64 " ns\n", CLOCKS[i].name, timespec_to_nanos(&out));
    }

    for (int i = 0; i < NUM_CLOCKS; i++)
    {
        printf("\n%s %zd samples:\n", CLOCKS[i].name, NUM_SAMPLES);
        int64_t samples[NUM_SAMPLES];

        for (int j = 0; j < NUM_SAMPLES; j++)
        {
            samples[j] = clock_nanos(CLOCKS[i].clock_id);
        }

        for (int j = 0; j < NUM_SAMPLES; j++)
        {
            printf("%" PRId64, samples[j]);
            if (j > 0)
            {
                int64_t diff = samples[j] - samples[j - 1];
                printf(" (diff=%" PRId64 ")", diff);
            }
            printf("\n");
        }
    }

    return 0;
}
