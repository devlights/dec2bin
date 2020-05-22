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
			name string
			in   testin
			out  testout
		}
	)

	cases := []testcase{
		{
			name: "convert 255",
			in:   testin{val: "255"},
			out:  testout{result: "0b11111111"},
		},
		{
			name: "convert 8",
			in:   testin{val: "8"},
			out:  testout{result: "0b1000"},
		},
		{
			name: "convert 4",
			in:   testin{val: "4"},
			out:  testout{result: "0b100"},
		},
		{
			name: "convert 2",
			in:   testin{val: "2"},
			out:  testout{result: "0b10"},
		},
		{
			name: "convert 1",
			in:   testin{val: "1"},
			out:  testout{result: "0b1"},
		},
		{
			name: "convert 15",
			in:   testin{val: "15"},
			out:  testout{result: "0b1111"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b, err := Convert(c.in.val)
			if err != nil {
				t.Error(err)
			}

			if c.out.result != b {
				t.Errorf("[want] %v\t[got] %v\n", c.out.result, b)
			}
		})
	}
}
