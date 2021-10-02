package process_test

import (
	"testing"

	"github.com/lmorg/murex/lang/process"
	"github.com/lmorg/murex/test/count"
)

func TestName(t *testing.T) {
	count.Tests(t, 4)

	name := new(process.Name)

	name.Set("foo")
	if name.String() != "foo" {
		t.Errorf("Set and/or String failed. Didn't return 'foo'")
	}

	name.Append("bar")
	if name.String() != "foobar" {
		t.Errorf("Set and/or String failed. Didn't return 'foobar'")
	}
}
