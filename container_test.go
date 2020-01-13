// Copyright © 2020 Hedzr Yeh.

package errorx_test

import (
	"github.com/hedzr/njuone/pkg/errorx"
	"io"
	"testing"
)

func sample(simulate bool) (err error) {
	c := errorx.NewContainer("sample error")
	if simulate {
		errorx.AttachTo(c, io.EOF, io.ErrUnexpectedEOF, io.ErrShortBuffer, io.ErrShortWrite)
	}
	err = c.Error()
	return
}

func TestContainer(t *testing.T) {
	err := sample(false)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%+v", err)
	}

	err = sample(true)
	if err == nil {
		t.Fatal("want error")
	} else {
		t.Logf("%+v", err)
	}
}
