package main

import (
	"fmt"

	vaderMicro "github.com/PIMPfiction/govader_backend"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	err := vaderMicro.Serve(e, "80")
	if err != nil {
		panic(err)
	}
	fmt.Scanln()

}
