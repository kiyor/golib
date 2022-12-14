package golib

import (
	"fmt"
)

type BasicTask struct {
	Frun           func() error
	FwhenError     func(error)
	FstopWhenError bool
}

func NewTask(run func() error, whenErr func(error), stopWhenError bool) *BasicTask {
	return &BasicTask{
		Frun:           run,
		FwhenError:     whenErr,
		FstopWhenError: stopWhenError,
	}
}

func (t *BasicTask) Run() error {
	if t.Frun != nil {
		return t.Frun()
	}
	return fmt.Errorf("function run cannot be nil")
}
func (t *BasicTask) WhenError(err error) {
	if t.FwhenError != nil {
		t.FwhenError(err)
	}
}
func (t *BasicTask) StopWhenError() bool {
	return t.FstopWhenError
}
