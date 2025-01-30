// 1. Given a string containing conversation between alice and bob,
// In the string, if it reaches $, it means end of alice message,
// and if it reaches #, it means end of bob's message,
// and if it reaches ^, it means end of conversation ignore string after that.

// Note: given string doesn't contain any spaces. If message contains two continuous conversations from one person it should be printed one after another as given in the example.

// write a program to separate out messages from alice and bob. Write messages from alice and bob on seperate channels, whenever a message from alice/bob, print it in front of their name as shown in the example below:

// Note: there is a single space before and after colon(:) and no space before and after comma.

// e.g: "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
// output: alice : helloBob,bob : helloalice,bob : howareyou?,alice : Iamgood.howareyou?

// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

package main

import "fmt"

func handleConversation(channelAlice chan string, channelBob chan string, conversation string) string {
	temp := ""
	finalStr := ""
	count := 0

	for _, char := range conversation {

		if string(char) == "$" {

			// received msg from Alice
			channelAlice <- temp
			temp = ""
			count += 1

		} else if string(char) == "#" {

			// received msg from Bob
			channelBob <- temp
			temp = ""
			count += 1

		} else if string(char) == "^" {
			break
		} else {
			temp = temp + string(char)
		}

		addMsgToFinalStr(channelAlice, channelBob, &finalStr, count)
	}

	return finalStr

}

func addMsgToFinalStr(channelAlice, channelBob chan string, finalStr *string, count int) {

	select {
	case m1 := <-channelAlice:
		// check if Alice's channel has a message, add it to our final String

		if count > 1 {
			*finalStr = *finalStr + ", " // to ensure `,` doesn't appear for first message
		}

		*finalStr = *finalStr + "alice : " + m1

	case m2 := <-channelBob:
		*finalStr = *finalStr + ", bob: " + m2 // check if Bob's channel has a message, add it to our final String

	default:

	}
}

func main() {

	conversation := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	channelAlice := make(chan string, 10)
	channelBob := make(chan string, 10)

	message := handleConversation(channelAlice, channelBob, conversation)

	fmt.Println(message)
}
