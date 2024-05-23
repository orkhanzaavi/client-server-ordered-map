package command_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"testwork/internal/command"
	"time"
)

func TestNewCommand(t *testing.T) {
	t.Run(
		"create a command successfully", func(t *testing.T) {
			cmd, err := command.NewCommand("test", "arg1", "arg2")

			t.Log("When the NewCommand function is called with the correct name and arguments")
			t.Log("Then the command should be created without error")
			assert.NoError(t, err)
			assert.Equal(t, "test", cmd.Name)
			assert.Equal(t, []string{"arg1", "arg2"}, cmd.Arguments)
		},
	)

	t.Run(
		"create a command with wrong name", func(t *testing.T) {
			cmd, err := command.NewCommand("testWrong", "arg1", "arg2")

			t.Log("When the NewCommand function is called with the incorrect name")
			t.Log("Then the command should not be created and an error should be returned")
			assert.ErrorIs(t, err, command.ErrInvalidCommandName)
			assert.Nil(t, cmd)
		},
	)

	t.Run(
		"create a command with the wrong arguments number", func(t *testing.T) {
			cmd, err := command.NewCommand("test", "arg1")

			t.Log("When the NewCommand function is called with the wrong number of arguments")
			t.Log("Then the command should not be created and an error should be returned")
			assert.ErrorIs(t, err, command.ErrInvalidNumberOfArguments)
			assert.Nil(t, cmd)
		},
	)
}

func TestFromString(t *testing.T) {
	t.Run(
		"create a command from string successfully", func(t *testing.T) {
			cmd, err := command.FromString("test:arg1,arg2")

			t.Log("When the FromString function is called with the correct string")
			t.Log("Then the command should be created without error")
			assert.NoError(t, err)
			assert.Equal(t, "test", cmd.Name)
			assert.Equal(t, []string{"arg1", "arg2"}, cmd.Arguments)
		},
	)

	t.Run(
		"create a command from string with wrong arguments", func(t *testing.T) {
			cmd, err := command.FromString("test")

			t.Log("When the FromString function is called with the wrong number of arguments")
			t.Log("Then the command should not be created and an error should be returned")
			assert.ErrorIs(t, err, command.ErrInvalidNumberOfArguments)
			assert.Nil(t, cmd)
		},
	)

	t.Run(
		"create a command with a colon in the end", func(t *testing.T) {
			cmd, err := command.FromString("getItem:")

			t.Log("When the command string has a colon in the end and number of arguments is 1")
			t.Log("Then the command should be created successfully")
			assert.NoError(t, err)
			assert.Equal(t, "getItem", cmd.Name)
			assert.Empty(t, cmd.Arguments[0])
		},
	)

	t.Run(
		"create a command without a colon in the end", func(t *testing.T) {
			cmd, err := command.FromString("getAllItems")

			t.Log("When the command string has no colon in the end and number of arguments is 0")
			t.Log("Then the command should be created successfully")
			assert.NoError(t, err)
			assert.Equal(t, "getAllItems", cmd.Name)
			assert.Empty(t, cmd.Arguments)
		},
	)
}

func TestReadCommandsFile(t *testing.T) {
	t.Run(
		"read commands from the file successfully", func(t *testing.T) {
			f, err := os.Create("/tmp/commands.txt")

			if err != nil {
				t.Fatal(err)
			}

			_, err = f.WriteString("test:arg1,arg2\n")

			if err != nil {
				t.Fatal(err)
			}
			err = f.Close()
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove("/tmp/commands.txt")

			ch, err := command.ReadCommandsFile("/tmp/commands.txt")

			var cmd *command.Command

			select {
			case cmd = <-ch:
			case <-time.After(time.Second):
				t.Fatal("timeout")
			}

			t.Log("When the ReadCommandsFile function is called with the correct file name")
			t.Log("Then the channel should be created without error")
			require.NoError(t, err)
			assert.NotNil(t, ch)
			t.Log("The channel should contain the command read from the file")
			assert.Equal(t, "test", cmd.Name)
			assert.Equal(t, []string{"arg1", "arg2"}, cmd.Arguments)
		},
	)

	t.Run(
		"read commands from the file with the wrong name", func(t *testing.T) {
			ch, err := command.ReadCommandsFile("/tmp/wrong-filename.txt")

			t.Log("When the ReadCommandsFile function is called with the wrong file name")
			t.Log("Then the channel should not be created and an error should be returned")
			assert.Error(t, err)
			assert.Nil(t, ch)
		},
	)
}
