package command

import (
	"bufio"
	"errors"
	"log/slog"
	"os"
	"strings"
)

// availableCommands is a map of available commands and the number of arguments they require
var availableCommands = map[string]int{
	"test":        2,
	"getItem":     1,
	"addItem":     2,
	"deleteItem":  1,
	"getAllItems": 0,
}

var ErrInvalidCommandName = errors.New("invalid command name")
var ErrInvalidNumberOfArguments = errors.New("invalid number of arguments")

type Command struct {
	Name      string
	Arguments []string
}

func NewCommand(name string, args ...string) (*Command, error) {
	availableCommand, ok := availableCommands[name]
	if !ok {
		return nil, ErrInvalidCommandName
	}
	if len(args) != availableCommand {
		return nil, ErrInvalidNumberOfArguments
	}
	return &Command{
		Name:      name,
		Arguments: args,
	}, nil
}

func FromString(str string) (*Command, error) {
	parts := strings.Split(str, ":")
	if len(parts) == 0 {
		return nil, ErrInvalidCommandName
	}
	name := parts[0]
	if len(parts) == 1 {
		return NewCommand(name)
	}
	args := strings.Split(parts[1], ",")

	return NewCommand(name, args...)
}

func ReadCommandsFile(fileName string) (chan *Command, error) {
	ch := make(chan *Command, 100)
	file, err := os.Open(fileName)

	if err != nil {
		//nolint:errcheck
		defer file.Close()
		slog.Error(
			"failed to open file",
			slog.String("err", err.Error()),
			slog.String("file", fileName),
		)
		return nil, err
	}
	r := bufio.NewReader(file)

	go func() {
		//nolint:errcheck
		defer file.Close()
		for {
			line, _, err := r.ReadLine()
			slog.Debug("read line", slog.String("line", string(line)))
			if len(line) > 0 {
				cmd, errConvert := FromString(string(line))
				if errConvert != nil {
					slog.Error(
						"failed to convert string to command",
						slog.String("err", errConvert.Error()),
						slog.String("line", string(line)),
					)
				}
				if cmd != nil {
					ch <- cmd
				}
			}
			if err != nil {
				close(ch)
				break
			}
		}
	}()

	return ch, nil
}

func (c *Command) String() string {
	if len(c.Arguments) == 0 {
		return c.Name
	}
	return c.Name + ":" + strings.Join(c.Arguments, ",")
}
