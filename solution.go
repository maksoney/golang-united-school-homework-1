package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	world_map := emoji.Sprint("Hello :world_map:!")
	fmt.Println(world_map)
}
