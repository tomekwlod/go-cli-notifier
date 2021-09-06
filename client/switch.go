package client

type BackendHTTPClient interface {
}

type Switch struct {
	client        BackendHTTPClient
	backendAPIURL string
	commands      map[string]func() func(string) error
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
		"help":   s.help,
	}

	return Switch{}
}
