package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/suzuken/dummy"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "dummy"
	app.Usage = "make a dummy data"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "format, f",
			Value: "string|10",
			Usage: "generate data type: like 'str|10,int|5,str|3'",
		},
		cli.IntFlag{
			Name: "length, l",
			Value: 1,
			Usage: "data record length",
		},
	}
	app.Action = func(c *cli.Context) {
		g := dummy.NewGenerator()
		s, err := g.Gen(c.String("format"), c.Int("length"))
		if err != nil {
			fmt.Printf("[ERR] %s", err)
		}
		fmt.Print(s)
		return
	}
	app.Run(os.Args)
}
