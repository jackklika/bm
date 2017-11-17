package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/exec"
)

func main() {
	bm := cli.NewApp()
	bm.Name = "bm"
	bm.Usage = "a bookmark utility for cd aliasing"

	bm.Action = func(c *cli.Context) error {

		path := c.Args().Get(0)

		stat, err := os.Stat(path)
		if os.IsNotExist(err) {
			fmt.Printf("%s is not a valid file or directory.\n", path)
			return nil
		} else if err != nil {
			fmt.Printf("Error parsing path.\n")
			return nil
		}

		var cmd *exec.Cmd

		if stat.Mode().IsDir() {
			fmt.Printf("%s\n", path)
			cmd = exec.Command("cd", path)
		} else if stat.Mode().IsRegular() {
			fmt.Printf("%s\n", path)
			cmd = exec.Command("vim", path)
		}

		er := cmd.Run()
		fmt.Printf("Error: %v", er)

		return nil
	}
	bm.Run(os.Args)
}
