package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(Tyrand(3))
	} else {
		fa := args[0]
		i64, err := strconv.ParseInt(fa, 10, 32)
		if err != nil {
			panic(err)
		}
		for i := 0; i < int(i64); i++ {
			fmt.Println(Tyrand(3))
		}
	}
}

const AllLetters = "abcdefghijklmnopqrstuvwxyz"
const AllNumbers = "0123456789"

// Letters with an underhanging element.
const UnderhangLetters = "gpqy"

// Too tall.
const TallLetters = "bdfhklt"

// Too similar to each other.
const DottedLetters = "ij"

// Zero can be confused for Capital Letter O
// 1 can be confused for Lowercase Letter l
// 6 and 9 can be confused for each other
const ExcludedNumbers = "0169"

const IncludedCharacters = "acemnorsuvwxz"
const IncludedNumbers = "234578"

// n is the number of couplets
func Tyrand(n int) (tyrand string) {
	for i := 0; i < n-1; i++ {
		couplet := Couplet()
		tyrand += string(couplet[:]) + "-"
	}
	final := Couplet()
	return tyrand + string(final[:])
}

func Couplet() (bytes [4]byte) {
	p1 := Pair()
	p2 := Pair()
	return [4]byte{p1[0], p1[1], p2[0], p2[1]}
}

func Pair() [2]byte {
	l := rand.Intn(len(IncludedCharacters))
	n := rand.Intn(len(IncludedNumbers))
	return [2]byte{IncludedCharacters[l], IncludedNumbers[n]}
}
