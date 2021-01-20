package runner

import "os/exec"

// Runner describes a basic command running interface
type Runner interface {
	Run(name string, arg ...string) ([]byte, error)
}

// ExecRunner runs commands using os.exec
type ExecRunner struct{}

// Run runs a command and returns byte slice output and a possible error
func (e *ExecRunner) Run(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).Output()
}
