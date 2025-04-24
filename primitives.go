package main

import "fmt"

func main() {
	//You create a channel of type string.
	//Channels are used to send data between go routines
	myChannel := make(chan string) //This is unBuffred channel
	anotherChannel := make(chan string)

	//This is how we will initialize buffered channels.
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}
	//This is Anonymous function are functions without a name
	//Basically this is sending data into channel
	go func() {
		myChannel <- "data"
	}() //We are calling it immediately by using ()

	go func() {
		anotherChannel <- "data2"
	}()
	//Main goroutine receives data from the channel
	// data := <-myChannel
	// fmt.Println(data)

	//As soon as any one of the channels is ready to send/receive, it executes the corresponding case.
	//if multiple channels are ready at the same time, one is chosen randomly.
	select {
	case firstChannel := <-myChannel:
		fmt.Println(firstChannel)
	case secondChannel := <-anotherChannel:
		fmt.Println(secondChannel)
	}

	for _, s := range chars {
		select {
		case charChannel <- s:

		}
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

}
