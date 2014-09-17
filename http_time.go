/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : httpTime.go

* Purpose :

* Creation Date : 09-17-2014

* Last Modified : Wed 17 Sep 2014 11:25:36 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package golib

import (
	"time"
)

var httpTimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

func Time2httpTime(t time.Time) string {
	return t.Format(httpTimeFormat)
}

func HttpTime2time(s string) (time.Time, error) {
	return time.Parse(httpTimeFormat, s)
}
