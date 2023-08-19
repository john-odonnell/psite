package main

import (
	"embed"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/john-odonnell/psite/v2/pkg/web"
)

//go:embed all:static
var staticDir embed.FS

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Printf("Unable to read envvar PORT: %s\n", err)
		fmt.Println("Falling back to port :3000")
		port = 3000
	}

	err = web.NewServer(staticDir, port).Listen()
	if err != nil {
		fmt.Printf("Error running server: %s\n", err)
	} else if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed.")
	}
}
