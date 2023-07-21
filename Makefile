CFLAGS=-Wall -Werror -O2 -g

all: clock_gettime_demo/clock_gettime_demo

clean:
	$(RM) clock_gettime_demo/clock_gettime_demo
	$(RM) -r clock_gettime_demo/clock_gettime_demo.dSYM
