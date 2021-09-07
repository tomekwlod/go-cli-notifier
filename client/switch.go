package client

import (
	"fmt"
	"os"
)

type BackendHTTPClient interface {
}

func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)

	s := Switch{
		client:        httpClient,
		backendAPIURL: uri,
	}

	s.commands = map[string]func() func(string) error{
		"create": s.create,
		"edit":   s.edit,
		"fatch":  s.fetch,
		"delete": s.delete,
		"health": s.health,
	}

	return s
}

type Switch struct {
	client        BackendHTTPClient
	backendAPIURL string
	commands      map[string]func() func(string) error
}

func (s *Switch) Switch() error {
	cmdName := os.Args[1]

	cmd, ok := s.commands[cmdName]

	if !ok {
		return fmt.Errorf("invalid command '%s'\n", cmdName)
	}

	return cmd()(cmdName)
}

func (s *Switch) Help() {
	var help string
	for name := range s.commands {
		help += name + "\t --help\n"
	}
	fmt.Printf("Usage of: %s:\n <command> [<args>]\n%s", os.Args[0], help)
}

func (s *Switch) create() func(string) error {
	return func(cmd string) error {
		fmt.Println("Create reminder")
		return nil
	}
}
func (s *Switch) edit() func(string) error {
	return func(cmd string) error {
		fmt.Println("Edit reminder")
		return nil
	}
}
func (s *Switch) fetch() func(string) error {
	return func(cmd string) error {
		fmt.Println("Fetch reminder")
		return nil
	}
}
func (s *Switch) delete() func(string) error {
	return func(cmd string) error {
		fmt.Println("Delete reminder")
		return nil
	}
}
func (s *Switch) health() func(string) error {
	return func(cmd string) error {
		fmt.Println("Health")
		return nil
	}
}
