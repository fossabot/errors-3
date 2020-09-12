// Copyright © 2020 Hedzr Yeh.

package errors

import (
	errorss "errors"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"testing"
)

func geneof() error {
	return io.EOF
}

func geneof2() error {
	return errors.Wrap(io.EOF, "xx")
}

func geneof13() error {
	return fmt.Errorf("xxx %w wrapped at go1.%v+", io.EOF, 13)
}

func geneofx() error {
	return WithCause(io.EOF, "text")
}

func geneofxs() error {
	return Wrap(io.EOF, "text")
}

func Test00(t *testing.T) {
	err := New("1")
	err = New("hello %v", "world")
	t.Log(err)
}

func Test01(t *testing.T) {
	err := WithCause(io.EOF, "1")
	err = WithCause(io.EOF, "hello %v", "world")
	t.Log(err)
	if !Is(err, io.EOF) {
		t.Fatal("is failed")
	}
}

func Test02(t *testing.T) {
	err := &withCause{
		causer: io.EOF,
		msg:    "dskjdl",
	}
	if !Is(err, io.EOF) {
		t.Fatal("is failed")
	}
}

func Test03(t *testing.T) {
	be := &bizErr{num: 2}
	err := &WithCauses{
		causers: []error{io.EOF, be},
		msg:     "dsda",
		Stack:   nil,
	}
	if !err.Is(io.EOF) {
		t.Fatal("is failed")
	}
	if !err.Is(be) {
		t.Fatal("is failed")
	}

	var e2 *bizErr
	if !err.As(&e2) {
		t.Fatal("as failed")
	}

	t.Log(err.Cause())
	t.Log(err.Causes())

	_ = err.SetCause(nil)
	_ = err.SetCause(io.EOF)

	err = &WithCauses{
		causers: nil,
		msg:     "dsda",
		Stack:   nil,
	}
	t.Log(err.Cause())
	t.Log(err.Causes())

	_ = err.SetCause(nil)
	_ = err.SetCause(io.EOF)
	_ = err.Unwrap()
}

func Test1(t *testing.T) {
	var err error

	err = geneof()
	if errors.Cause(err) == io.EOF {
		t.Logf("ok: %v", err)
	} else {
		t.Fatal("expect it is a EOF")
	}

	err = geneof2()
	if errors.Cause(err) == io.EOF {
		t.Logf("ok: %v", err)
	} else {
		t.Fatal("expect it is a EOF")
	}

	err = geneof13()
	if errorss.Is(err, io.EOF) {
		t.Logf("ok: %v", err)
	} else {
		t.Fatal("expect it is a EOF")
	}
}

func Test2(t *testing.T) {
	var err error

	err = geneofx()
	if errors.Cause(err) == io.EOF {
		t.Logf("ok: %v", err)
	} else {
		t.Fatal("expect it is a EOF")
	}
	if Cause(err) == io.EOF {
		t.Logf("ok: %+v", err)
	} else {
		t.Fatal("expect it is a EOF")
	}

	err = geneofxs()
	if errors.Cause(err) == io.EOF {
		t.Logf("ok: %v", err)
	} else {
		t.Fatal("expect it is a EOF")
	}

	// errorx tests -------------------------------

	// errorx.Cause() and Cause1()
	if Cause(err) == io.EOF {
		// Wrap(err, msg): the error object has stacktrace info
		t.Logf("ok: %+v", err)
	} else {
		t.Fatal("Cause() failed: expect it is a EOF")
	}
	if Is(err, io.EOF) {
		// Wrap(err, msg): the error object has stacktrace info
		t.Logf("ok: %+v", err)
	} else {
		t.Fatal("Is() failed: expect it is a EOF")
	}
	if Unwrap(Unwrap(err)) == io.EOF {
		// Wrap(err, msg): the error object has stacktrace info
		t.Logf("ok: %+v", err)
	} else {
		t.Fatal("Unwrap() failed: expect it is a EOF")
	}

	var perr *os.PathError
	err = Wrap(&os.PathError{Err: io.EOF, Op: "find", Path: "/"}, "wrong path and rights")
	if As(err, &perr) {
		t.Logf("ok: %+v", *perr)
	} else {
		t.Fatal("As() failed: expect it is a os.PathError{}")
	}

	// var c = NewContainer("container")
	// AttachTo(c, io.EOF, io.ErrShortBuffer, io.ErrUnexpectedEOF)
	// t.Logf("ok: %+v | container is empty: %v", c, ContainerIsEmpty(c))
	// if ContainerIsEmpty(c) != false {
	// 	t.Fatal("ContainerIsEmpty(c) failed: expect it is false.")
	// }
}
