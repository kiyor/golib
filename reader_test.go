package golib_test

import (
	"fmt"
	"github.com/kiyor/golib"
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
	golib.LineReaderFunc(rd, p)
}
func TestLineReaderChan(t *testing.T) {
	rd, e := os.Open("testfile")
	if e != nil {
		t.Fatal(e)
	}
	done := make(chan struct{})
	ch := make(chan string)
	err := golib.LineReaderChan(rd, ch, done)
	go func() {
		select {
		case e := <-err:
			t.Fatal(e)
		}

	}()
	go func() {
		for l := range ch {
			fmt.Println(l)
		}
	}()
	<-done
}
