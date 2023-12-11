package barpear

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"
)

func Main() int {
	slog.Debug("barpear", "test", true)

	test()
	return 0
}

type GeneratorOption func(*rand.Rand)

// WithSeed is a functional option to set the seed for random number generation.
func WithSeed(seed int64) GeneratorOption {
	return func(rng *rand.Rand) {
		rng.Seed(seed)
	}
}

func RandomPositiveIntegerSliceUpToMax(max int, options ...GeneratorOption) []int {
	randomSlice := make([]int, max+1)

	for i := 0; i <= max; i++ {
		randomSlice[i] = i
	}

	// Default seed with current time for randomness
	randSrc := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSrc)

	// Apply functional options
	for _, option := range options {
		option(rng)
	}

	// Shuffle the slice to get a random order
	for i := len(randomSlice) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		randomSlice[i], randomSlice[j] = randomSlice[j], randomSlice[i]
	}

	return randomSlice
}

func test() {
	max := 10
	// Without specifying seed (uses current time)
	slice1 := RandomPositiveIntegerSliceUpToMax(max)
	fmt.Println("slice 1:", slice1)

	// With a specific seed for reproducibility
	slice2 := RandomPositiveIntegerSliceUpToMax(max, WithSeed(42))
	fmt.Println("slice 2:", slice2)
}
