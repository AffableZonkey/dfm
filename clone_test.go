package dfm_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/chasinglogic/dfm"
	"github.com/chasinglogic/dfm/test"
)

func TestCreateLongURL(t *testing.T) {
	url, user := dfm.CreateURL(strings.Split("https://bitbucket.org/chasinglogic/dotfiles", "/"))
	if url != "https://bitbucket.org/chasinglogic/dotfiles" {
		t.Errorf("Expected: %s Got: %s", "https://bitbucket.org/chasinglogic/dotfiles", url)
	}

	if user != "chasinglogic" {
		t.Errorf("Expected: %s Got: %s", "chasinglogic", user)
	}
}

func TestCreateShortURL(t *testing.T) {
	url, user := dfm.CreateURL(strings.Split("chasinglogic/dotfiles", "/"))
	if url != "https://github.com/chasinglogic/dotfiles" {
		t.Errorf("Expected: %s Got: %s", "https://github.com/chasinglogic/dotfiles", url)
	}

	if user != "chasinglogic" {
		t.Errorf("Expected: %s Got: %s", "chasinglogic", user)
	}
}

func TestCreateSSHURL(t *testing.T) {
	url, user := dfm.CreateURL(strings.Split("git@github.com:/chasinglogic/dotfiles", "/"))
	if url != "git@github.com:/chasinglogic/dotfiles" {
		t.Errorf("Expected: %s Got: %s", "git@github.com:/chasinglogic/dotfiles", url)
	}

	if user != "chasinglogic" {
		t.Errorf("Expected: %s Got: %s", "chasinglogic", user)
	}
}

func TestCloneRepo(t *testing.T) {
	tmp, e := test.InitTest(false)
	if e != nil {
		t.Error(e)
	}

	e = dfm.CloneRepo("https://github.com/chasinglogic/dotfiles",
		filepath.Join(tmp.TmpCfg, "chasinglogic"))
	if e != nil {
		t.Error(e)
	}

	testFilesExistence(filepath.Join(tmp.TmpCfg, "chasinglogic"))
}
