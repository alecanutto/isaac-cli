package internal

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

type Command interface {
	Name() string
	Example() string
	Help() string
	LongHelp() string
	Register(*flag.FlagSet)
	Run()
}

type CommandRoot struct {
	Name     string
	commands []Command
}

func CommandInit(name string) *CommandRoot {
	return &CommandRoot{
		Name: name,
	}
}

func (cr *CommandRoot) Start(commandList []Command) error {
	if len(commandList) == 0 {
		return errors.New("Obrigatório no mínimo um comando para ser executado")
	}
	cr.commands = commandList
	if len(os.Args) < 2 {
		cr.ShowHelp()
		return errors.New("Por favor, informe um comando")
	}

	userCommand := ArgumentFilter(os.Args[1:])
	if userCommand.Command == "" {
		cr.ShowHelp()
		return errors.New("Por favor, informe um comando válido")
	}
	if userCommand.Command == "help" {
		return nil
	}
	for _, command := range cr.commands {
		if userCommand.Command == command.Name() {
			fs := flag.NewFlagSet(command.Name(), flag.ExitOnError)
			command.Register(fs)
			fs.Parse(os.Args[2:])
			command.Run()
			return nil
		}
	}
	cr.ShowHelp()
	return errors.New("Comando não encontrado")
}

func (cr *CommandRoot) ShowHelp() {
	fmt.Printf("Usage: %s [COMMAND] [OPTIONS]\n\n", cr.Name)

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "commands:\n\n")

	for _, command := range cr.commands {
		fmt.Fprintf(tw, "\t- %s\t%s\n", command.Name(), command.Help())
	}

	tw.Flush()

	fmt.Fprintf(tw, "\n\nexamples:\n\n")

	for _, command := range cr.commands {
		fmt.Fprintf(tw, "\t%s\n", command.Example())
	}

}
