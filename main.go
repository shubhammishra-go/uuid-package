package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	str := uuid.NewString()

	fmt.Println(uuid.Parse("23118942-af95-41b2-ad59-c4b6970c29c"))

	fmt.Println(str)
}
