package client

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type BackendHTTPClient interface {
	Create(title, message string, duration time.Duration) ([]byte, error)
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

		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(createCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(createCmd); err != nil {
			return err
		}

		res, err := s.client.Create(*t, *m, *d)
		if err != nil {
			return wrapError("Coulnd not create reminder", err)
		}

		fmt.Printf("Reminder created:\n%s", string(res))

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

func (s *Switch) reminderFlags(f *flag.FlagSet) (*string, *string, *time.Duration) {
	t, m, d := "", "", time.Duration(0)

	f.StringVar(&t, "title", "", "Reminder title")
	f.StringVar(&t, "t", "", "Reminder title")
	f.StringVar(&m, "message", "", "Reminder message")
	f.StringVar(&m, "m", "", "Reminder message")
	f.DurationVar(&d, "duration", 0, "Reminder duration")
	f.DurationVar(&d, "d", 0, "Reminder duration")

	return &t, &m, &d
}
func (s *Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])

	if err != nil {
		return wrapError("Could not parse '"+cmd.Name()+"' command flags", err)
	}
	return nil
}
func (s *Switch) checkArgs(minArgs int) error {

	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}

	if len(os.Args)-2 < minArgs {
		fmt.Printf("Incorrect use of %s\n%s %s --help\n", os.Args[1], os.Args[0], os.Args[1])
		return fmt.Errorf("%s expects at least %d arg(s), %d provided", os.Args[1], minArgs, len(os.Args)-2)
	}

	return nil
}
