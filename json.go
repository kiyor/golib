/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : json.go

* Purpose :

* Creation Date : 02-02-2017

* Last Modified : Mon 15 May 2017 01:34:10 AM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package golib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go4.org/errorutil"
	"os"
)

func JsonUnmarshal(data []byte, v interface{}) error {
	fh := bytes.NewReader(data)
	d := json.NewDecoder(fh)
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		extra := ""
		if serr, ok := err.(*json.SyntaxError); ok {
			if _, serr := fh.Seek(0, os.SEEK_SET); serr != nil {
				return fmt.Errorf("seek error: %v", err)
			}
			line, col, highlight := errorutil.HighlightBytePosition(fh, serr.Offset)
			extra = fmt.Sprintf(":\nError at line %d, column %d (file offset %d):\n%s",
				line, col, serr.Offset, highlight)
		}
		return fmt.Errorf("error parsing JSON object %s\n%v",
			extra, err)
	}
	return nil
}
