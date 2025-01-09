package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpperCase(t *testing.T) {
	input := "Hello"
	expected := "HELLO"
	result := UpperCase(input)
	assert.Equal(t, expected, result, "expected %s, got %s", expected, result)
}

func TestLowerCase(t *testing.T) {
	input := "Hello"
	expected := "hello"
	result := LowerCase(input)

	if expected != result {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestConcatenation(t *testing.T) {
	input := "Hello "
	concat := "World!"
	expected := "Hello World!"
	result := Concat(input, concat)
	assert.Equal(t, expected, result)
}

func TestSlicing(t *testing.T) {
	input := "I will crash this tech interview!"
	start := 18
	end := len(input)

	expected := "tech interview!"
	result := Slicing(input, start, end)

	assert.Equal(t, expected, result, "expected %s, got %s", expected, result)
}

func TestContains(t *testing.T) {
	input := "Hello, World!"
	sub := "World"
	expected := true
	result := Contains(input, sub)
	assert.Equal(t, expected, result, "expected %v, got %v", expected, result)
}

func TestReplace(t *testing.T) {
	input := "Hello, World!"
	old := "World"
	new := "Gopher"
	expected := "Hello, Gopher!"
	result := Replace(input, old, new)
	assert.Equal(t, expected, result, "expected %s, got %s", expected, result)
}

func TestHasPrefix(t *testing.T) {
	input := "Hello, World!"
	prefix := "Hello"
	expected := true
	result := HasPrefix(input, prefix)
	assert.Equal(t, expected, result, "expected %v, got %v", expected, result)
}

func TestHasSuffix(t *testing.T) {
	input := "Hello, World!"
	suffix := "World!"
	expected := true
	result := HasSuffix(input, suffix)
	assert.Equal(t, expected, result, "expected %v, got %v", expected, result)
}

func TestSplit(t *testing.T) {
	input := "a,b,c"
	sep := ","
	expected := []string{"a", "b", "c"}
	result := Split(input, sep)
	assert.Equal(t, expected, result, "expected %v, got %v", expected, result)
}

func TestJoin(t *testing.T) {
	elements := []string{"a", "b", "c"}
	sep := "-"
	expected := "a-b-c"
	result := Join(elements, sep)
	assert.Equal(t, expected, result, "expected %s, got %s", expected, result)
}

func TestTrim(t *testing.T) {
	input := "***Hello***"
	cutset := "*"
	expected := "Hello"
	result := Trim(input, cutset)
	assert.Equal(t, expected, result, "expected %s, got %s", expected, result)
}

func TestTrimSpace(t *testing.T) {
	input := "   Hello   "
	expected := "Hello"
	result := TrimSpace(input)
	assert.Equal(t, expected, result, "expected %s, got %s", expected, result)
}

func TestRepeat(t *testing.T) {
	input := "Hello"
	count := 3
	expected := "HelloHelloHello"
	result := Repeat(input, count)
	assert.Equal(t, expected, result, "expected %s, got %s", expected, result)
}

func TestFields(t *testing.T) {
	input := "Hello World"
	expected := []string{"Hello", "World"}
	result := Fields(input)
	assert.Equal(t, expected, result, "expected %v, got %v", expected, result)
}

func TestCount(t *testing.T) {
	input := "banana"
	sub := "a"
	expected := 3
	result := Count(input, sub)
	assert.Equal(t, expected, result, "expected %d, got %d", expected, result)
}

func TestCompare(t *testing.T) {
	a := "Hello"
	b := "World"
	expected := -1
	result := Compare(a, b)
	assert.Equal(t, expected, result, "expected %d, got %d", expected, result)
}

func TestEqualFold(t *testing.T) {
	a := "Hello"
	b := "hello"
	expected := true
	result := EqualFold(a, b)
	assert.Equal(t, expected, result, "expected %v, got %v", expected, result)
}
