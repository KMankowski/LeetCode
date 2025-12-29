package main

import "fmt"

func main() {
	fmt.Println(longestPrefixNaive("ababab"))
}

func longestPrefixNaive(s string) string {
	longestPrefix := ""
	for i := range len(s) - 1 {
		currPrefix := s[:i+1]
		currSuffix := s[len(s)-i-1:]
		if currPrefix == currSuffix {
			longestPrefix = currPrefix
		}
	}
	return longestPrefix
}
