package main

import "fmt"

func main() {
	fmt.Println(say("bot", "hello!"))
}

func say(user, message string) string {
	return fmt.Sprintf("I am %s, %s", user, message)
}
