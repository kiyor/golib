/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : osexec.go

* Purpose :

* Creation Date : 12-16-2014

* Last Modified : Tue 16 Dec 2014 08:57:37 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package golib

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"strings"
)

func Osexec(cmd string) (stdOut string, stdErr error) {
	c := exec.Command("/bin/sh", "-c", cmd)

	stdout, err := c.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return "", err
	}

	outStdCh := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, stdout)
		outStdCh <- buf.String()
	}()

	outErrCh := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, stderr)
		outErrCh <- buf.String()
	}()

	err = c.Start()
	if err != nil {
		return "", err
	}
	defer c.Wait()
	out := <-outStdCh
	e := <-outErrCh
	out = strings.Trim(out, "\n")
	e = strings.Trim(e, "\n")
	if len(e) > 0 {
		return out, errors.New(e)
	}
	return out, nil
}
