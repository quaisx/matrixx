package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
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

	fmt.Printf("%s\n", r)
	fmt.Println(strings.Repeat("-", 50))
}
