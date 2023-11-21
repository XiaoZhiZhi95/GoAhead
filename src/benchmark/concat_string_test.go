package benchmark

import (
	"bytes"
	"testing"
)

func ConcatStringByAdd(elems []string) string {
	ret := ""
	for _, elem := range elems {
		ret += elem
	}

	return ret
}

func ConcatStringByBytesBuffer(elems []string) string {
	var buf bytes.Buffer
	for _, elem := range elems {
		buf.WriteString(elem)
	}

	return buf.String()
}
/*
func TestConcatStringByAdd(t *testing.T) {
	a := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ConcatStringByAdd(elems)
	a.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	a := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ConcatStringByBytesBuffer(elems)
	a.Equal("12345", ret)
}
*/
func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.StartTimer()
	for i:=0;i<b.N;i++ {
		ConcatStringByAdd(elems)
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.StartTimer()
	for i:=0;i<b.N;i++ {
		ConcatStringByBytesBuffer(elems)
	}
	b.StopTimer()
}