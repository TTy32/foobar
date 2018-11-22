package main

import (
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
    "fmt"
)

type Payload struct {
	ID   string
    Data types.JSONText
}

func (p Payload) Validate() error {
	if p.ID != "theID" {
		return errors.New("Invalid ID")
	}
    return nil
}


func main() {
    fmt.Println("Hello")
}

