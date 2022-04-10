package main

import (
	"fmt"
	"os"

	vaderMicro "github.com/PIMPfiction/govader_backend"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	err := vaderMicro.Serve(e, os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	fmt.Scanln()

}
