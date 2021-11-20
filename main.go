package main

import (
	"djinn/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome the Djinn Programming langugae!\n", user.Username)
	fmt.Printf("Feel free to mess around and see what things you can uncover\n")
	repl.Start(os.Stdin, os.Stdout)
}
