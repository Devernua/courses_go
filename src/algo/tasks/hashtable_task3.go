package main

import "fmt"

const primeHashTable uint64 = 1000000007 // prime
const randomX uint64 = 263
const obrX uint64 = 836501907


func calcHash(pwrs []uint64, str *string) (result uint64) {
	for i := 0; i < len(*str); i++ {
		newVal := result + pwrs[i] * uint64((*str)[i])
		result = uint64(newVal % primeHashTable)
	}

	return result
}

func main() {
	var pattern string
	var text string
	fmt.Scanf("%v", &pattern)
	fmt.Scanf("%v", &text)

	var pwrs []uint64
	pwrs = append(pwrs, 1)
	for i := 1; i < len(pattern); i++ {
		newVal := randomX * pwrs[i - 1]
		pwrs = append(pwrs, newVal % primeHashTable)
	}

	patternHash := calcHash(pwrs, &pattern)

	curTextSlice := text[:len(pattern)]
	curHash := calcHash(pwrs, &curTextSlice)

	if curHash == patternHash && pattern == curTextSlice {
		fmt.Printf("%v ", 0)
	}

	for i := 1; i + len(pattern) <= len(text); i++ {
		tP := uint64(text[i + len(pattern) - 1]) * pwrs[len(pattern) - 1] % primeHashTable
		curHash = ((curHash + primeHashTable - uint64(curTextSlice[0])) * obrX + tP) % primeHashTable
		curTextSlice = text[i:len(pattern) + i]
		if curHash == patternHash && pattern == curTextSlice {
			fmt.Printf("%v ", i)
		}
	}

	fmt.Println()
}
