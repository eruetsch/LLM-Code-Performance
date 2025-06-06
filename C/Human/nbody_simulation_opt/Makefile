# C Flags, -g, -Ofast

CFLAGS := -Wall -pedantic -std=c17 -Ofast -march=native -mavx2 -fomit-frame-pointer -fopenmp -ffast-math -falign-functions=64
CC := gcc
LDLIBS := -lm -fopenmp
INC :=

# run make -j 8, seems to be fastest

# ENDING is cpp

BUILD_DIR := build
SRCS := src/kdtree.c src/particle.c
#$(shell find src -iname '*.c' -type f)
HEADERS := $(shell find src -iname '*.h' -type f)
OBJS := $(SRCS:src/%.c=$(BUILD_DIR)/%.o)
OUT := kdtree-sim

MAIN_OBJ := $(BUILD_DIR)/kdtree-sim.o
TEST_OBJ := $(BUILD_DIR)/test.o
BENCH_OBJ := $(BUILD_DIR)/benchmark.o

.PHONY: clean lines run benchmark test


$(OUT): $(OBJS) $(MAIN_OBJ)
	$(CC) $(CFLAGS) -o $@ $^ $(LDLIBS)

# run-test: test
# 	./$<

test: $(OBJS) $(TEST_OBJ)
	$(CC) $(CFLAGS) -o $@ $^ $(LDLIBS)

benchmark: $(OBJS) $(BENCH_OBJ)
	$(CC) $(CFLAGS) -o $@ $^ $(LDLIBS)

$(BUILD_DIR)/%.o: src/%.c $(HEADERS)
	@mkdir -p $(@D)
	$(CC) -c $(INC) -o $@ $< $(CFLAGS)

clean:
	rm -rf $(BUILD_DIR) $(OUT) test benchmark

lines: $(SRCS) $(HEADERS)
	wc -l $^ | grep total

# Run with different thread counts to find optimal settings
run-scaling: $(OUT)
	@echo "Running scaling tests with different thread counts..."
	@for t in 1 2 4 8 16; do \
		echo "\n=== Using $$t threads ==="; \
		OMP_NUM_THREADS=$$t ./$(OUT) 10 10000; \
	done
