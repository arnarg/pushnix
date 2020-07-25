package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GitPush(r string, f bool) error {
	args := []string{"push", r}
	if f {
		args = append(args, "-f")
	}
	return runCommand("git", args...)
}

func SSHNixosRebuild(h *SSHHost, u bool, e []string) error {
	remoteCommand := "sudo nixos-rebuild switch"
	if u {
		remoteCommand += " --upgrade"
	}
	if len(e) > 0 {
		remoteCommand = fmt.Sprintf("%s %s", remoteCommand, strings.Join(e, " "))
	}
	return runCommand("ssh", "-t", "-l", h.User, h.Host, remoteCommand)
}

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
