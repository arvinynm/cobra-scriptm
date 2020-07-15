package cmd

import (
	"fmt"
	"log"
	"os/user"
	"testing"
)

func TestHome(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HomeDir:", u.HomeDir)
}
