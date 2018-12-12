package golib

import (
	"fmt"
	"os"
	"testing"
)

func TestLineReaderFunc(t *testing.T) {
	rd, err := os.Open("testfile")
	if err != nil {
		t.Fatal(err)
	}
	p := func(l string) error {
		fmt.Println(l)
		return nil
	}
	LineReaderFunc(rd, p)
}
func TestLineReaderChan(t *testing.T) {
	rd, e := os.Open("testfile")
	if e != nil {
		t.Fatal(e)
	}
	ch := make(chan string)
	err := LineReaderChan(rd, ch)
	go func() {
		select {
		case e := <-err:
			t.Fatal(e)
		}

	}()
	for l := range ch {
		fmt.Println(l)
	}
}
