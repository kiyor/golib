/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : naming.go

* Purpose :

* Creation Date : 08-06-2014

* Last Modified : Wed Aug  6 12:31:24 2014

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package golib

import (
	"fmt"
	"strconv"
	"strings"
)

func ToNameAuto(name string) string {
	if strings.ToUpper(name[:1]) == name[:1] {
		return ToUnderscoreName(name)
	} else {
		return ToCamelName(name)
	}
}

func ToUnderscoreName(name string) string {
	var res string
	for k, v := range name {
		s := fmt.Sprintf("%c", v)
		if strings.ToUpper(s) == s {
			if _, err := strconv.Atoi(s); err != nil && k != 0 {
				res += "_"
			}
			s = strings.ToLower(s)
		}
		res += s
	}
	return res
}

func ToCamelName(name string) string {
	part := strings.Split(name, "_")
	var res string
	for _, v := range part {
		v = strings.ToUpper(v[:1]) + v[1:]
		res += v
	}
	return res
}
