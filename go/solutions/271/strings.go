package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func main() {
	s := new(Solution)
	encoded := s.Encode([]string{"Hello", "World"})
	fmt.Println(s.Decode(encoded))
}

func (s *Solution) Encode(strs []string) string {
	var encodedBuilder strings.Builder
	for _, str := range strs {
		encodedBuilder.WriteString(strconv.Itoa(len(str)) + "." + str)
	}
	return encodedBuilder.String()
}

func (s *Solution) Decode(encoded string) []string {
	var strs []string
	i := 0
	for i < len(encoded) {
		var length int
		length, i = parseLength([]rune(encoded), i)
		strs = append(strs, parseString([]rune(encoded), i, length))
		i += length
	}
	return strs
}

func parseString(s []rune, i, length int) string {
	return string(s[i : i+length])
}

func parseLength(s []rune, i int) (int, int) {
	rawLength := ""
	for s[i] != '.' {
		rawLength += string(s[i])
		i++
	}
	parsedLength, _ := strconv.Atoi(rawLength)
	return parsedLength, i + 1
}
