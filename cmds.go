package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// CmdsInterface is an interface for available go-serverless-cli commands
type CmdsInterface interface {
	Execute()
}

type cmds struct {
	args []string
}

// NewCmds is used to initialized commands class
func NewCmds(args []string) CmdsInterface {
	return &cmds{
		args: args,
	}
}

// Execute runs the command based on passed arguments
func (c *cmds) Execute() {
	if len(c.args) <= 1 {
		c.help()
		return
	}

	switch c.args[1] {
	default:
		c.help()
	case `setup`:
		c.setup()
	case `create`:
		c.create()
	case `build`:
		c.build()
	}

}

func (c *cmds) help() {
	fmt.Println(`go-serverless-cli help`)
}

func (c *cmds) setup() {
	filepath := GOPATH + `/src/github.com/madebyais/go-serverless-cli/install.sh`

	output, err := exec.Command(`/bin/sh`, filepath).CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}

	fmt.Println(`GO-SERVERLESS-CLI dependencies have been installed successfully.`)
}

func (c *cmds) create() {
	pluginName := c.args[2]
	filename := fmt.Sprintf(`%s.go`, pluginName)
	filepath := fmt.Sprintf(`./%s`, filename)

	data := strings.Replace(DEFAULT_FUNCTION_FILE, `{plugin_name}`, pluginName, -1)
	bytefile := []byte(data)
	err := ioutil.WriteFile(filepath, bytefile, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(fmt.Sprintf(`%s has been created`, filename))
}

func (c *cmds) build() {
	if len(c.args) < 3 {
		fmt.Println(`Please define plugin name`)
		return
	}

	pluginName := c.args[2]
	buildcmd := fmt.Sprintf(`./%s.so ./%s.go`, pluginName, pluginName)

	output, err := exec.Command(`go`, `build`, `-buildmode=plugin`, `-o`, buildcmd).CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}

	fmt.Println(fmt.Sprintf(`%s has been build`, pluginName))
}
