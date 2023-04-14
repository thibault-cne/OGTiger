package options

import "fmt"

type Options struct {
	AST  string
	File string

	Steps int
}

func Parse(osArgs []string) (*Options, error) {
	var options Options
	options.Steps = 1

	flags, err := parse(osArgs)

	if err != nil {
		return nil, err
	}

	for _, flag := range flags {
		switch flag.Flag {
		case &AST:
			options.AST = flag.Value
			options.Steps++
		case &File:
			options.File = flag.Value
			options.Steps++
		}
	}

	if options.File == "" {
		return nil, fmt.Errorf("no file specified")
	}

	return &options, nil
}