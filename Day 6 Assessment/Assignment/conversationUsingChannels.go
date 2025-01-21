package main

import (
	"fmt"
	"sync"
)

func main() {

	conversation := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	var wg sync.WaitGroup

	channelAlice := make(chan string, 10)
	channelBob := make(chan string, 10)

	temp := ""

	finalStr := ""

	for i := 0; i < len(conversation); i++ {

		if conversation[i:i+1] == "$" {
			// received msg from Alice
			wg.Add(1)

			// add msg to Alice's channel
			go addMsgToAlice(channelAlice, temp, &wg)
			temp = ""
			wg.Wait()

		} else if conversation[i:i+1] == "#" {
			// received msg from Bob
			wg.Add(1)

			// add message to Bob's channel
			go addMsgToBob(channelBob, temp, &wg)
			temp = ""
			wg.Wait()

		} else if conversation[i:i+1] == "^" {
			break
		} else {
			temp = temp + conversation[i:i+1]
		}

		// using Select to check which channel has received a msg...
		// if msg received, add it to conversation

		select {
		case m1 := <-channelAlice:
			finalStr = finalStr + ", alice : " + m1 // check if Alice's channel has a message, add it to our final String
		case m2 := <-channelBob:
			finalStr = finalStr + ", bob: " + m2 // check if Alice's channel has a message, add it to our final String
		default:
			continue
		}
	}
	fmt.Println(finalStr)

}

func addMsgToAlice(alice chan string, msg string, wg *sync.WaitGroup) {
	alice <- msg
	wg.Done()
}

func addMsgToBob(bob chan string, msg string, wg *sync.WaitGroup) {
	bob <- msg
	wg.Done()
}
