package repository

import "testing"

type getValueTest struct {
	arg1     string
	expected int
}

var getValueTests = []getValueTest{
	getValueTest{"glob", 1},
	getValueTest{"prok", 5},
	getValueTest{"pish", 10},
	getValueTest{"tegj", 50},
}

func TestGetValue(t *testing.T) {

	for _, test := range getValueTests {
		if output := GetValue(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
