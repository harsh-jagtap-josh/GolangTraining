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

func handleConversation(channelAlice chan string, channelBob chan string, conversation string) string {

	temp := ""

	finalStr := ""

	for i, _ := range conversation {

		if conversation[i:i+1] == "$" {

			// received msg from Alice
			channelAlice <- temp
			temp = ""

		} else if conversation[i:i+1] == "#" {

			// received msg from Bob
			channelBob <- temp
			temp = ""

		} else if conversation[i:i+1] == "^" {
			break

		} else {
			temp = temp + conversation[i:i+1]
		}

	}

	addMsgToFinalStr(channelAlice, channelBob, finalStr)

	return finalStr

}

func addMsgToFinalStr(channelAlice, channelBob chan string, finalStr string) {

	select {
	case m1 := <-channelAlice:
		finalStr = finalStr + ", alice : " + m1 // check if Alice's channel has a message, add it to our final String

	case m2 := <-channelBob:
		finalStr = finalStr + ", bob: " + m2 // check if Alice's channel has a message, add it to our final String

	default:

	}
}

func main() {

	conversation := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	channelAlice := make(chan string, 10)
	channelBob := make(chan string, 10)

	handleConversation(channelAlice, channelBob, conversation)

}
