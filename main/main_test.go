package main

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	fmt.Println("func init1")
}

func init() {
	fmt.Println("func init2")
}

func TestMain(m *testing.M) {
	fmt.Println("before testing")
	code := m.Run()
	fmt.Println("after testing")
	os.Exit(code)
}

func Test(t *testing.T) {
}
