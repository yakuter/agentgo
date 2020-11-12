package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var usageStr = `
Usage: agentgo-server [options]
Server Options:
    -app <application>          Application to execute command
    -arg <argument>             Arguments for application

Common Options:
    -help                       Show this message
`

func PrintUsageErrorAndDie(err error) {
	log.Println(err)
	fmt.Println(usageStr)
	os.Exit(1)
}

func PrintHelpAndDie() {
	fmt.Println(usageStr)
	os.Exit(0)
}

// Options is main value holder agentgo-server flags.
type Options struct {
	App      string     `json:"app"`
	Args     arrayFlags `json:"args"`
	ShowHelp bool       `json:"show_help"`
}

// ConfigureOptions accepts a flag set and augments it with agentgo-server
// specific flags. On success, an options structure is returned configured
// based on the selected flags.
func ConfigureOptions(fs *flag.FlagSet, args []string) (*Options, error) {

	// Create empty options
	opts := &Options{}

	// Define flags
	fs.BoolVar(&opts.ShowHelp, "help", false, "Show help message")
	fs.StringVar(&opts.App, "app", "", "Application to execute command: ls")
	fs.Var(&opts.Args, "arg", "Arguments for application: -lah")

	// Parse arguments and check for errors
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	// If it is not help and app is empty, return error
	if opts.ShowHelp == false && opts.App == "" {
		err := errors.New("application argument (-app <application>) is empty")
		return nil, err
	}

	return opts, nil
}
