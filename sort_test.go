package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestSort(t *testing.T) {

	sigs := make(chan os.Signal, 1)
	defer close(sigs)
	done := make(chan bool, 1)
	defer close(done)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func(sigs <-chan os.Signal, done <-chan bool) {
		fmt.Println("Signal routine is running")
		for {
			select {
			case s := <-sigs:
				fmt.Println("Signal:", s)
				return
			case <-done:
				fmt.Println("Done")
				return
			}
		}
	}(sigs, done)

	tt := time.Now()
	fmt.Printf(" >>> Running TestSort: %s\n", tt.Format(time.ANSIC))
	fmt.Println(strings.Repeat("-", 50))

	n1 := new(Node[int])
	n1.value = 1

	n2 := new(Node[int])
	n2.value = 2
	n1.next = n2

	n3 := new(Node[int])
	n3.value = 3
	n2.next = n3

	n4 := new(Node[int])
	n4.value = 4
	n3.next = n4

	n5 := new(Node[int])
	n5.value = 5
	n4.next = n5

	fmt.Printf("%s\n", n1)

	r := n1.Reverse()
	rs := fmt.Sprintf("%s", r)
	if rs != string("5, 4, 3, 2, 1") {
		t.Errorf("Incorrect reverse sequence: %s", r)
	}
	done <- true
	fmt.Printf("%s\n", r)
	fmt.Println(strings.Repeat("-", 50))
}
