package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-git/go-git/v5"
)

type SSHHost struct {
	User string
	Host string
	Port string
}

func GetHostFromRemoteName(n string) (*SSHHost, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	gitDir, err := findGitBase(wd)
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(gitDir)
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

func findGitBase(p string) (string, error) {
	if p == "/" {
		return "", fmt.Errorf("not a git repository")
	}
	if fm, err := os.Stat(p + "/.git"); err == nil && fm.IsDir() {
		return p, nil
	}
	ret, err := findGitBase(filepath.Dir(p))
	return ret, err
}
