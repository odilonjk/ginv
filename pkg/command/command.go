package command

// Cmd is the command to be executed
type Cmd interface {
	// Execute runs command
	Execute()
}
