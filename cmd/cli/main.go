package main

import (
	"fmt"

	"github.com/alecanutto/gonboarding/command/start"
	"github.com/alecanutto/gonboarding/command/text"
	"github.com/alecanutto/gonboarding/internal"
)

func main() {
	list := []internal.Command{
		new(start.Start),
		new(text.Text),
	}

	if err := internal.CommandInit("isaacli").Start(list); err != nil {
		fmt.Println(err)
	}
}
