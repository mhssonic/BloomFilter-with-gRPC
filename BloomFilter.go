package main

import (
	"fmt"
	"hash/fnv"
)

const m uint64 = 350 * 1e7

var bitArray = make([]bool, m)

func main() {
	added := []string{"hey", "ali", "pashmaki", "sfds", "sdidd", "reza", "hi", "dude", "nun", "tof"}
	notAdded := []string{"hy", "hadi", "shokolat", "wqeds", "hgfde", "hasan", "james", "uwuwu", "12312f", "12edw"}
	for _, str := range added {
		addString(str)
	}
	for _, str := range added {
		fmt.Printf("%s %t\n", str, BloomFilterIsThere(str))
	}
	fmt.Printf("-------------------\n")
	for _, str := range notAdded {
		fmt.Printf("%s %t\n", str, BloomFilterIsThere(str))
	}

}

func addString(str string) {
	i, j, k := hash(str)
	bitArray[i] = true
	bitArray[j] = true
	bitArray[k] = true
}

func BloomFilterIsThere(str string) bool {
	i, j, k := hash(str)
	return bitArray[i] && bitArray[j] && bitArray[k]
}

// return 3 hash code of a string
func hash(s string) (uint64, uint64, uint64) {
	h := fnv.New64a()
	h.Write([]byte(s))
	a := h.Sum64() % m
	h.Write([]byte(s))
	b := h.Sum64() % m
	h.Write([]byte(s))
	c := h.Sum64() % m
	return a, b, c
}
