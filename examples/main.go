package main

import (
	"fmt"
	"log/slog"

	"github.com/andriusm/testlib"
)

func main() {
	config := testlib.Config{
		Greeting: "Hello",
		Logger:   slog.Default(),
	}

	tl, err := testlib.New(config)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	result := tl.DoThings("World", 2)

	fmt.Println(result)
}
