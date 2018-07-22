package xutils_test

import (
	"github.com/dimiro1/x/xutils"
	"os"
	"testing"
)

func Test_GetenvDefault(t *testing.T) {
	type testCase struct {
		Before   func(testCase)
		Key      string
		Expected string
		Default  string
	}

	cases := []testCase{
		{
			Before: func(t testCase) {
				os.Clearenv()
				os.Setenv(t.Key, t.Expected)
			},
			Key:      "SOME_KEY",
			Expected: "HELLO",
			Default:  "",
		},
		{
			Before: func(t testCase) {
				os.Clearenv()
				os.Setenv(t.Key, t.Expected)
			},
			Key:      "SOME_KEY",
			Expected: "",
			Default:  "",
		},
		{
			Before: func(t testCase) {
				os.Clearenv()
			},
			Key:      "SOME_KEY",
			Expected: "DEFAULT_VALUE",
			Default:  "DEFAULT_VALUE",
		},
	}

	for _, test := range cases {
		t.Run(test.Key, func(t *testing.T) {
			test.Before(test)

			got := xutils.GetenvDefault(test.Key, test.Default)
			if got != test.Expected {
				t.Errorf("expected %s, got %s", test.Expected, got)
			}
		})
	}
}
