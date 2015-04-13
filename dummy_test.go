package dummy_test

import (
	"github.com/suzuken/dummy"
	"os"
	"strings"
	"testing"
)

// global dummy generator, test purpose only
var g = &dummy.Generator{}

func TestMain(m *testing.M) {
	g = dummy.NewGenerator()
	os.Exit(m.Run())
}

func TestCheckLengthString(t *testing.T) {
	l := g.String(10)
	if len(l) != 10 {
		t.Fatalf("string length is not proper: %s", l)
	}
	i := g.Int(20)
	if len(i) != 20 {
		t.Fatalf("int length is not proper: %s", i)
	}
}

func TestGenLine(t *testing.T) {
	f := "str|3,str|10,int|10,str|3"
	s := g.GenLine(f)
	ss := strings.Split(s, ",")
	if len(ss) != 4 {
		t.Fatalf("generate string is failed, length is wrong %v %v", f, ss)
	}
}

func TestFormat(t *testing.T) {
	f1 := "str|3,str|10,int|10,str|3"
	if !dummy.IsProperFormat(f1) {
		t.Fatalf("%s is proper format, but not detect", f1)
	}
	f2 := "a|3"
	if dummy.IsProperFormat(f2) {
		t.Fatalf("%s is invalid format, but not detect", f2)
	}
	f3 := "str|||||10"
	if dummy.IsProperFormat(f3) {
		t.Fatalf("%s is invalid format, but not detect", f3)
	}
	f4 := "str10"
	if dummy.IsProperFormat(f4) {
		t.Fatalf("%s is invalid format, but not detect", f4)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g.String(100)
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g.Int(100)
	}
}
