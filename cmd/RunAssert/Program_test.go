package main

import (
	"fmt"
	"testing"
)

func TestEnroll(t *testing.T) {
	fmt.Println(Enroll("127.0.0.1:88", ":89", "喵", "喵"))
}
