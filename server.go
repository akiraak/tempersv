package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os/exec"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		cmd := exec.Command("date")
		cmdoutput, err := cmd.Output()
		output := string(cmdoutput)
		if err != nil {
			output = fmt.Sprintf("%s\n%s", output, err)
		}
		fmt.Println(output)

		return c.String(http.StatusOK, output)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
