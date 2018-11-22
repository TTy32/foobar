package main

import (
	"testing"
)

func TestValidation(t *testing.T) {
    payload := Payload{
        ID: "theID",
        Data: nil,
    }

    err := payload.Validate()

    if err != nil {
        t.Errorf("Failed validation: %v", err)
    }
}




