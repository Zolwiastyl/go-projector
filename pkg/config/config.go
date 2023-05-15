package projector_config

import (
	"fmt"
	"os"
	"path"
)

type Operation = int

const (
	Print Operation = iota
	Add
	Remove
)

type Config struct {
	Args      []string
	Operation Operation

	Pwd        string
	ConfigPath string
}

func NewConfig(options *Options) (*Config, error) {
	pwd, err := getPwd(options)
	if err != nil {
		return nil, err
	}

	configPath, err := getConfigPath(options)
	if err != nil {
		return nil, err
	}

	args, err := getArgs(options)
	if err != nil {
		return nil, err
	}
	operation := getOperation(options)

	return &Config{
		Args:       args,
		Pwd:        pwd,
		ConfigPath: configPath,
		Operation:  operation,
	}, nil
}

func getPwd(options *Options) (string, error) {
	if options.Pwd != "" {
		return options.Pwd, nil
	}
	return os.Getwd()
}

func getConfigPath(options *Options) (string, error) {
	if options.Config != "" {
		return options.Config, nil
	}
	config, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(config, ".projector.json"), nil

}
func getOperation(options *Options) Operation {

	if len(options.Args) == 0 {
		return Print
	}
	if options.Args[0] == "add" {
		return Add
	}
	if options.Args[0] == "rm" {
		return Remove
	}
	return Print
}
func getArgs(options *Options) ([]string, error) {
	numberOfArgumentsProvided := len(options.Args) - 1
	operation := getOperation(options)
	if len(options.Args) == 0 {
		return []string{}, nil
	}
	if operation == Add {
		if len(options.Args) != 3 {
			return nil, fmt.Errorf("add command requires 2 arguments but got %v", numberOfArgumentsProvided)
		}
		return options.Args[1:], nil
	}
	if operation == Remove {
		if len(options.Args) != 2 {
			return nil, fmt.Errorf("remove command requires 1 arguments but got %v", numberOfArgumentsProvided)
		}
		return options.Args[1:], nil
	}
	if operation == Print {
		if len(options.Args) > 1 {
			return nil, fmt.Errorf("print requires 0 or one arguments, but got %v", numberOfArgumentsProvided)
		}
	}
	return options.Args, nil
}
