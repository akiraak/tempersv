package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		cmd := exec.Command("temper-poll")
		cmdoutput, err := cmd.Output()
		output := string(cmdoutput)

		if err == nil {
			lines := strings.Split(string(cmdoutput), "\n")
			temps := strings.Split(lines[1], " ")
			celsius := temps[2]
			fahrenheit := temps[3]
			output = fmt.Sprintf("akiraak宅の室温\n摂氏:%s 華氏:%s", celsius, fahrenheit)
		} else {
			output = err.Error()
		}
		fmt.Println(output)
		return c.String(http.StatusOK, output)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
