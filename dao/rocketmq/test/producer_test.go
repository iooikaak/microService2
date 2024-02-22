package test

import (
	"fmt"
	"testing"
)

func TestSendToTest(t *testing.T) {
	msg := "testtttttttttttttttts"
	err := d.SendToTest(ctx, msg, "", "test")
	if err != nil {
		fmt.Printf("Producer Send err: %v", err)
		return
	}
}
