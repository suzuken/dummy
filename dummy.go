// dummy is dummy data generator. This package provide the dummy data in specified format.
//
// The format in dummy is simple. If you want to generate 10 letter string, format is as below:
//
//     str|10
//
// Also, this support separated value.
//
//     str|5,str|10,int|5
//
// This create 5 letter string, 10 letter string and 5 letter integers in comma separated format.
package dummy

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyz")
	numbers = []rune("0123456789")
	any     = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
)

// Generator for dummy.
type Generator struct {
	rn *rand.Rand
}

// Initialize Generator
func NewGenerator() *Generator {
	return &Generator{
		rn: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// check if format is proper or not
func IsProperFormat(format string) bool {
	r := regexp.MustCompile(`(str|int)\|[0-9]+`)
	s := r.FindAllString(format, -1)
	return len(s) > 0
}

// Generate separated string in the specified format and line length.
func (g *Generator) Gen(format string, len int, separator string) (string, error) {
	if !IsProperFormat(format) {
		return "", errors.New("provided format is not proper")
	}
	var s string
	for i := 0; i < len; i++ {
		s = s + g.GenLine(format, separator) + "\n"
	}
	return s, nil
}

// Generate each line. You can use dummy format for generate line.
func (g *Generator) GenLine(format, separator string) string {
	ret := make([]string, 0)
	for _, f := range parseFormat(format) {
		k, v := parseFieldKV(f)
		switch {
		case k == "str":
			ret = append(ret, g.String(v))
		case k == "int":
			ret = append(ret, g.Int(v))
		}
	}
	return strings.Join(ret, separator)
}

func parseFieldKV(field string) (k string, v int) {
	kv := strings.Split(field, "|")
	k = kv[0]
	v, _ = strconv.Atoi(kv[1])
	return
}

func parseFormat(format string) []string {
	return strings.Split(format, ",")
}

// manipulate each character by specified runes.
func (g *Generator) GenRune(r []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = r[g.rn.Int()%len(r)]
	}
	return string(b)
}

// Generate random string in particular length.
func (g *Generator) String(n int) string {
	return g.GenRune(letters, n)
}

// Generate random integer in particular length.
func (g *Generator) Int(n int) string {
	return g.GenRune(numbers, n)
}
