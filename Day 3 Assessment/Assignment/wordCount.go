package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	maxx := 0

	words := strings.Split(line, " ")

	var freq = map[string]int{}

	for i := 0; i < len(words); i++ {
		freq[words[i]]++
		maxx = max(maxx, freq[words[i]])
	}

	ans := []string{}

	pre := map[string]bool{}

	for i := 0; i < len(words); i++ {
		if freq[words[i]] == maxx {
			if pre[words[i]] == false {
				ans = append(ans, words[i])
				pre[words[i]] = true
			}
		}
	}

	fmt.Println(ans)
}
