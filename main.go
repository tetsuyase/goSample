package main

import (
	"github.com/google/uuid"
)

func main() {
	var uuidList []uuid.UUID
	for i := 0; i < 4; i++ {
		uuidList = append(uuidList, uuid.New())
	}
}
