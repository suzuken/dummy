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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Value: "string|10",
			Usage: "generate data type: like 'str|10,int|5,str|3'",
		},
		cli.IntFlag{
			Name:  "length, l",
			Value: 1,
			Usage: "data record length",
		},
		cli.StringFlag{
			Name:  "separator, s",
			Value: ",",
			Usage: "field separator",
		},
	}
	app.Action = func(c *cli.Context) {
		g := dummy.NewGenerator()
		if !dummy.IsProperFormat(c.String("format")) {
			fmt.Fprint(os.Stderr, "provided format is invalid.")
			os.Exit(1)
			return
		}
		for i := 0; i < c.Int("length"); i++ {
			fmt.Println(g.GenLine(c.String("format"), c.String("separator")))
		}
		return
	}
	app.Run(os.Args)
}
