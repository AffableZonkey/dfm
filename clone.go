package dfm

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	cli "gopkg.in/urfave/cli.v1"
)

// Clone will clone the given git repo to the profiles directory, it optionally
// will call link or use depending on the flag given.
func Clone(c *cli.Context) error {
	url, user := CreateURL(strings.Split(c.Args().First(), "/"))
	userDir := filepath.Join(getProfileDir(), user)
	if cloneErr := CloneRepo(url, userDir); cloneErr != nil {
		return cloneErr
	}

	if c.Bool("link") {
		return Link(c)
	}

	return nil
}

// CreateURL will add the missing github.com for the shorthand version of
// links.
func CreateURL(s []string) (string, string) {
	if len(s) == 2 {
		return fmt.Sprintf("https://github.com/%s", strings.Join(s, "/")), s[0]
	}

	return strings.Join(s, "/"), s[len(s)-2]
}

// CloneRepo will git clone the provided url into the appropriate profileDir
func CloneRepo(url, profileDir string) error {
	if CONFIG.Verbose {
		fmt.Printf("Creating profile in %s\n", profileDir)
	}

	c := exec.Command("git", "clone", url, profileDir)
	_, err := c.CombinedOutput()
	if err != nil && err.Error() == "exit status 128" {
		return cli.NewExitError("Profile exists, perhaps you meant dfm update or link?", 128)
	}

	return err
}
