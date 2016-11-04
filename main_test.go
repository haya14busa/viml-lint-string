package main

import (
	"bytes"
	"io/ioutil"

	"github.com/kylelemons/godebug/diff"
)

import "testing"

func TestRun(t *testing.T) {
	okb, err := ioutil.ReadFile("testdata/test.ok")
	if err != nil {
		t.Fatal(err)
	}
	ok := string(okb)

	var buf bytes.Buffer

	if err := run(&buf, []string{"testdata/test.vim"}); err != nil {
		t.Error(err)
	}

	if d := diff.Diff(buf.String(), ok); d != "" {
		t.Errorf("dff: (-got, +want)\n%s", d)
	}
}
