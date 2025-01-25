package commands

type ICommands interface {
	Execute(message string) string
}

type CommandImpl struct{}

func (c *CommandImpl) Execute(message string) string {
	return "!PONG"
}

func CommandFactory(cmd string) ICommands {
	switch cmd {
	case "ping":
		return &CommandImpl{}
	default:
		return nil
	}

}
