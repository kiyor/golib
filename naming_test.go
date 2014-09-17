/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : naming_test.go

* Purpose :

* Creation Date : 08-26-2014

* Last Modified : Tue 26 Aug 2014 01:28:49 AM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package golib

import (
	"log"
	"testing"
)

func Test_ToNameAuto(t *testing.T) {
	name := "Test1String"
	name = ToNameAuto(name)
	log.Println(name)
	name = ToNameAuto(name)
	log.Println(name)
}
