package main

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
)

func main() {
	fmt.Printf("New UUID: %s", uuid.NewUUID())
}
