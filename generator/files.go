package generator

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// FileData defines fields needed by the generator for creating a file.
type FileData struct {
	Filename string
	Template string
}

// DirectoryData defines fields needed by the generator for creating a directory.
type DirectoryData struct {
	Dirname string
	Empty bool
}

// createFileContents substitutes the owner and projectName into the template string.
func createFileContents(template string, owner string, projectName string) (c string) {
	c = strings.ReplaceAll(template, "$1", owner)
	c = strings.ReplaceAll(c, "$2", projectName)
	return c
}

// GitAddFile executes the git add command.
func GitAddFile(filename string) error {
	if err := exec.Command("git", "add", filename).Run(); err != nil {
		return errors.WithMessagef(err, "failed to add %s to Git", filename)
	}

	return nil
}

// MustGitAddFile wraps GitAddFile and panics if GitAddFile returns an error.
func MustGitAddFile(filename string) {
	if err := GitAddFile(filename); err != nil {
		panic(errors.WithStack(err))
	}
}

// CreateAndAddFile creates file contents from the template, owner, and projectName then writes the file and adds it to
// Git.
func CreateAndAddFile(filename string, template string, owner string, projectName string) error {
	contents := createFileContents(template, owner, projectName)

	if err := ioutil.WriteFile(filename, []byte(contents), 0644); err != nil {
		return errors.WithMessagef(err, "failed to create file %s", filename)
	}

	if err := GitAddFile(filename); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// CreateAndAddDirectory creates a new directory. If the directory is flagged as empty, it will then create a .gitkeep
// file and add the directory to Git. Otherwise this function will assume that the directory will be added to Git later
// after it has been populated or that the user does not wish to add this directory to Git.
func CreateAndAddDirectory(dirname string, empty bool) error {
	if err := os.Mkdir(dirname, 0755); err != nil {
		return errors.WithMessagef(err, "failed to create directory %s", dirname)
	}

	if empty {
		f, err := os.Create(dirname+"/.gitkeep")
		if err != nil {
			return errors.WithMessagef(err, "failed to create .gitkeep for %s/", dirname)
		}
		if err := f.Close(); err != nil {
			return errors.WithMessagef(err, "failed to create .gitkeep for %s/", dirname)
		}
		if err := GitAddFile(dirname+"/.gitkeep"); err != nil {
			return errors.WithMessagef(err, "failed to add %s/ to Git", dirname)
		}
	}

	return nil
}

// CreateAndAddFileList iterates over the given FileData, executing CreateAndAddFile for each one.
func CreateAndAddFileList(files []FileData, owner string, projectName string) (e []error) {
	for _, v := range files {
		if err := CreateAndAddFile(v.Filename, v.Template, owner, projectName); err != nil {
			e = append(e, errors.WithStack(err))
		}
	}

	return e
}

// CreateAndAddDirectoryList iterates over the given DirectoryData, executing CreateAndAddDirectory for each one.
func CreateAndAddDirectoryList(dirs []DirectoryData) (e []error) {
	for _, v := range dirs {
		if err := CreateAndAddDirectory(v.Dirname, v.Empty); err != nil {
			e = append(e, errors.WithStack(err))
		}
	}

	return e
}
