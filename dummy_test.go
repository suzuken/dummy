package dummy_test

import (
	"github.com/suzuken/dummy"
	"testing"
)

func TestCheckLengthString(t *testing.T) {
	l := dummy.String(10)
	if len(l) != 10 {
		t.Fatalf("string length is not proper: %s", l)
	}
	i := dummy.Int(20)
	if len(i) != 20 {
		t.Fatalf("int length is not proper: %s", i)
	}
}
