package util

import (
	"fmt"
	"regexp"

	"github.com/go-git/go-git/v5"
)

type SSHHost struct {
	User string
	Host string
	Port string
}

func GetHostFromRemoteName(n string) (*SSHHost, error) {
	options := &git.PlainOpenOptions{
		DetectDotGit: true,
	}
	repo, err := git.PlainOpenWithOptions(".", options)
	if err != nil {
		return nil, err
	}

	remote, err := repo.Remote(n)
	if err != nil {
		return nil, err
	}

	return getHostFromRemote(remote)
}

func getHostFromRemote(r *git.Remote) (*SSHHost, error) {
	for _, u := range r.Config().URLs {
		host := parseGitURL(u)
		if host != nil {
			return host, nil
		}
	}
	return nil, fmt.Errorf("Could not find a URL that matches an SSH string in remote %s", r.Config().Name)
}

func parseGitURL(u string) *SSHHost {
	r := regexp.MustCompile(`^([^@]+)@([^:]+):?(\d*).*$`)
	m := r.FindStringSubmatch(u)
	if m == nil {
		return nil
	}
	user := m[1]
	host := m[2]
	port := m[3]
	if user == "" || host == "" {
		return nil
	}

	if port == "" {
		port = "22"
	}

	return &SSHHost{
		User: user,
		Host: host,
		Port: port,
	}
}
