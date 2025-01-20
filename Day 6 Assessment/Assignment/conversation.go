// have a doubt in How to use Channels in this question, so used a simple normal approach to solve this problem..

package main

import "fmt"

func main() {

	conversation := "helloBob$helloalice#howareyou?$Iamgood.howareyou?$^"

	temp := ""

	finalStr := ""

	alice := true

	count := 0
	for i := 0; i < len(conversation); i++ {
		if conversation[i:i+1] == "$" {
			if count > 0 {
				finalStr = finalStr + ","
			}
			if alice {
				finalStr = finalStr + "alice : "
			} else {
				finalStr = finalStr + "bob : "
			}
			count++
			finalStr = finalStr + temp
			temp = ""
			alice = !alice
		} else if conversation[i:i+1] == "#" {
			if !alice {
				finalStr = finalStr + ",bob : "
			} else {
				finalStr = finalStr + ",alice : "
			}
			finalStr = finalStr + temp
			temp = ""
		} else if conversation[i:i+1] == "^" {
			break
		} else {
			temp = temp + conversation[i:i+1]
		}
	}
	fmt.Println(finalStr)
}
