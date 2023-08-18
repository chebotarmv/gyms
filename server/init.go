package server

import (
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var PackSizes []int

func init() {
	var err error
	packSizesStr := os.Getenv("PACK_SIZES")
	PackSizes, err = parsePackSizes(packSizesStr)
	if err != nil {
		log.Fatal(err)
	}
}

func parsePackSizes(packSizesStr string) ([]int, error) {
	if packSizesStr == "" {
		m := []int{5000, 2000, 1000, 500, 250}
		return m, nil
	}
	slice := make([]int, 0)
	sizes := strings.Split(packSizesStr, ",")
	for _, sizeStr := range sizes {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			return nil, errors.New("invalid input")
		}
		slice = append(slice, size)
	}
	// remove duplicates
	m := make(map[int]bool)
	for _, size := range slice {
		m[size] = true
	}
	packSizesSlice := make([]int, 0)
	for size := range m {
		packSizesSlice = append(packSizesSlice, size)
	}

	// sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(packSizesSlice)))
	return packSizesSlice, nil
}
