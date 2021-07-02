package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	fmt.Println("~~~Queue~~~")

	JQ := NewJobQueue()

	err := JQ.JoinIntoQueue(NewJob(time.Second, simplePrint("hola")))
	if err != nil {
		fmt.Println("The queue is full")
	}
	err = JQ.JoinIntoQueue(NewJob(time.Second, fillForm))
	if err != nil {
		fmt.Println("The queue is full")
	}
	err = JQ.JoinIntoQueue(NewJob(time.Second, timeFunc))
	if err != nil {
		fmt.Println("The queue is full")
	}

	fmt.Println(JQ)
	time.Sleep(time.Hour)
}

func simplePrint(message string) func() {
	return func() {
		fmt.Println(message)
	}
}

func fillForm() {
	var name, country string
	fmt.Print("Name: ")
	// Scan no funciona en testing pero en main.go si
	fmt.Scan(&name)
	fmt.Print("County: ")
	fmt.Scan(&country)
	fmt.Scanln()
	fmt.Printf("Name: %s | Country: %s\n", name, country)

}

func timeFunc() {
	fmt.Println("Wait 3 Seconds")
	time.Sleep(time.Second * 3)
	fmt.Println("Finish")
}
