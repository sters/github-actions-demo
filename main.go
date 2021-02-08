package main

import "fmt"

func main() {
	fmt.Println(say("hello!"))
}

func say(message string) string {
	return fmt.Sprintf("I'm bot, %s", message)
}
