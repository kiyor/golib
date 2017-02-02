/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : json.go

* Purpose :

* Creation Date : 02-02-2017

* Last Modified : Thu 02 Feb 2017 06:54:19 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package golib

import (
	"bytes"
	"encoding/json"
)

func JsonUnmarshal(data []byte, v interface{}) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(&v)
}
