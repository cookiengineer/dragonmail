package main

import (
	"dragonmail/structs"
	"fmt"
)

func main() {

	profile := structs.NewProfile("~/Dragon")

	fmt.Println(profile)

}
