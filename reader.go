package golib

import (
	"bufio"
	"io"
	"strings"
)

func LineReaderFunc(rd io.Reader, action func(string) error) error {
	reader := bufio.NewReader(rd)
	for {
		l, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		} else {
			err := action(strings.TrimRight(l, "\n"))
			if err != nil {
				return err
			}
		}
	}
}
func LineReaderChan(rd io.Reader, ch chan string, done chan struct{}) chan error {
	err := make(chan error)
	go func() {
		reader := bufio.NewReader(rd)
		for {
			l, e := reader.ReadString('\n')

			if e != nil {
				close(ch)
				if e != io.EOF {
					err <- e
				}
				done <- struct{}{}
				return
			} else {
				ch <- strings.TrimRight(l, "\n")
			}
		}
	}()
	return err
}
