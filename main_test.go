package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	type (
		testin struct {
			val string
		}
		testout struct {
			result string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in:  testin{val: "255"},
			out: testout{result: "0b11111111"},
		},
		{
			in:  testin{val: "8"},
			out: testout{result: "0b1000"},
		},
		{
			in:  testin{val: "4"},
			out: testout{result: "0b100"},
		},
		{
			in:  testin{val: "2"},
			out: testout{result: "0b10"},
		},
		{
			in:  testin{val: "1"},
			out: testout{result: "0b1"},
		},
		{
			in:  testin{val: "15"},
			out: testout{result: "0b1111"},
		},
	}

	for _, c := range cases {
		b, err := Convert(c.in.val)
		if err != nil {
			t.Error(err)
		}

		if c.out.result != b {
			t.Errorf("[want] %v\t[got] %v\n", c.out.result, b)
		}
	}
}
