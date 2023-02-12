package main

import (
	"testing"

	"github.com/emerfauzan/cake-store-api/test"
)

func TestMain(t *testing.T) {
	test := test.Init()
	test.AuthLogin(t)
	test.GetCakes(t)
}
