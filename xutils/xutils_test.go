package xutils_test

import (
	"os"
	"testing"

	"github.com/dimiro1/x/xutils"
)

func Test_GetenvDefault(t *testing.T) {
	//noinspection GoUnusedType
	type scenario struct {
		Before   func(scenario)
		Key      string
		Expected string
		Default  string
	}

	scenarios := []scenario{
		{
			Before: func(s scenario) {
				os.Clearenv()
				os.Setenv(s.Key, s.Expected)
			},
			Key:      "SOME_KEY",
			Expected: "HELLO",
			Default:  "",
		},
		{
			Before: func(s scenario) {
				os.Clearenv()
				os.Setenv(s.Key, s.Expected)
			},
			Key:      "SOME_KEY",
			Expected: "",
			Default:  "",
		},
		{
			Before: func(s scenario) {
				os.Clearenv()
			},
			Key:      "SOME_KEY",
			Expected: "DEFAULT_VALUE",
			Default:  "DEFAULT_VALUE",
		},
	}

	for _, aScenario := range scenarios {
		t.Run(aScenario.Key, func(t *testing.T) {
			aScenario.Before(aScenario)

			got := xutils.GetenvDefault(aScenario.Key, aScenario.Default)
			if got != aScenario.Expected {
				t.Errorf("expected %s, got %s", aScenario.Expected, got)
			}
		})
	}
}
