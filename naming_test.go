/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : naming_test.go

* Purpose :

* Creation Date : 08-26-2014

* Last Modified : Tue 16 Dec 2014 07:11:56 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package golib

import (
	// 	"log"
	"testing"
)

func Test_ToNameAuto(t *testing.T) {
	name := "Test1String"
	name = ToNameAuto(name)
	if name != "test1_string" {
		t.Fatal("failed ToNameAuto")
	}
	t.Log(name)
	name = ToNameAuto(name)
	if name != "Test1String" {
		t.Fatal("failed ToNameAuto")
	}
	t.Log(name)
}
