package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arrAnagram := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	output := make(map[string][]string)
	for i := 0; i < len(arrAnagram); i++ {
		splitI := strings.Split(arrAnagram[i], "")
		sort.Strings(splitI)
		sortedSplitI := strings.Join(splitI, "")
		output[sortedSplitI] = append(output[sortedSplitI], arrAnagram[i])
	}
	for key, value := range output {
		sort.Strings(value)
		fmt.Println("Anagram", key, "Value is :", value)
	}
}
