package main

import (
	"os"
)

func main() {
	cmds := NewCmds(os.Args)
	cmds.Execute()
}
