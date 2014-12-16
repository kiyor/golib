/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : osexec_test.go

* Purpose :

* Creation Date : 12-16-2014

* Last Modified : Tue 16 Dec 2014 07:08:57 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package golib

import (
	//	"fmt"
	"log"
	"testing"
)

func TestExec(t *testing.T) {
	_, err := Osexec("cat osexec_test.go")
	if err != nil {
		log.Println(err.Error())
	} else {
		t.Log("osexec pass")
	}
	_, err = Osexec("1")
	if err != nil {
		if err.Error() == "/bin/sh: 1: command not found" {
			t.Log("osexec pass")
		}
	}
}
