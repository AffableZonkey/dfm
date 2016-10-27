// Package test provides various utilites only used in unit tests.
package test

import (
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

// FilesExistence tests all file paths sent to it determining if they exist
// returning any errors it encounters.
func FilesExistence(files ...string) error {
	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			return err
		}
	}

	return nil
}

// Dirs just aggregates all the tmp directories for a given test.
type Dirs struct {
	Tmp    string
	TmpHm  string
	TmpCfg string
}

// InitTest will create temporary directories that are unique to each test.
func InitTest(createProfile bool) (*Dirs, error) {
	tmp := filepath.Join(os.TempDir(), "dfm_test"+string(rand.Int()))
	t := &Dirs{
		tmp,
		filepath.Join(tmp, "home"),
		filepath.Join(tmp, "config"),
	}

	e := os.Mkdir(t.Tmp, os.ModeDir)
	if e != nil {
		return t, e
	}

	e = os.Mkdir(t.TmpHm, os.ModeDir)
	if e != nil {
		return t, e
	}

	e = os.Mkdir(t.TmpCfg, os.ModeDir)
	if e != nil {
		return t, e
	}

	return t, nil
}

// Cleanup cleans up the tmp dirs after the test.
func Cleanup(ts *testing.T, t *Dirs) {
	err := os.RemoveAll(t.Tmp)
	if err != nil {
		ts.Errorf("Error cleaning up: %s\n", err.Error())
	}

	err = os.RemoveAll(t.TmpHm)
	if err != nil {
		ts.Errorf("Error cleaning up: %s\n", err.Error())
	}

	err = os.RemoveAll(t.TmpCfg)
	if err != nil {
		ts.Errorf("Error cleaning up: %s\n", err.Error())
	}
}
