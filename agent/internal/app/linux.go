package app

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Linux struct {
	App  string
	Args []string
}

func (l *Linux) Execute() (string, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(l.App, l.Args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		return outStr, fmt.Errorf(errStr)
	}

	return outStr, nil
}
