package migration

import "os/exec"

func poMigrate(option Option) error {
	if option.Po.NotRun {
		return nil
	}

	command := exec.Command(option.Po.Gormt())
	return command.Run()
}
