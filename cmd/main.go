// It provides teh coral project sponge CLI  platform.
package main

import (
	"fmt"
	"os"

	"github.com/ardanlabs/kit/cfg"
	"github.com/coralproject/coral-talk-driver/pkg/talk"
)

// export TALK_URL="http://localhost:3000/api/v1/"
const (
	cfgTALKURL = "TALK_URL"
)

func main() {

	cfgTALKURL, err := cfg.String(cfgTALKURL)

	if err := talk.Connect(cfgTALKURL); err != nil {
		fmt.Printf("main: Unable to initialize configuration. %v", err)
		os.Exit(-1)
	}

	id := 1
	if result, err := talk.GetComment(id); err != nil {
		fmt.Printf("main: Unable to execute the command. %v", err)
		os.Exit(-1)
	}

	fmt.Print(result)
}
