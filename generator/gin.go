package generator

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"os/user"
)

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(errors.WithMessagef(err, "failed to execute %s", command))
	}
}

func NewGinProject(owner string, projectName string) {
	u, err := user.Current()
	if err != nil {
		panic(errors.WithMessage(err, "unable to get current user"))
	}

	projDir := u.HomeDir+"/devel/prog/"+projectName
	if err := os.Mkdir(projDir, 0755); err != nil {
		panic(errors.WithMessage(err, "failed to create project directory"))
	}
	if err := os.Chdir(projDir); err != nil {
		panic(errors.WithMessage(err, "unable to enter project directory"))
	}
	runCommand("git", "init")

	dirs := []DirectoryData{
		{
			Dirname: "config",
			Empty:   false,
		},
		{
			Dirname: "controllers",
			Empty:   false,
		},
		{
			Dirname: "db",
			Empty:   false,
		},
		{
			Dirname: "db/migrations",
			Empty:   true,
		},
		{
			Dirname: "db/seed",
			Empty:   false,
		},
		{
			Dirname: "models",
			Empty:   false,
		},
	}
	if err := CreateAndAddDirectoryList(dirs); err != nil {
		for _, e := range err {
			fmt.Println(e)
		}
		panic("errors encountered while creating directories")
	}

	files := []FileData{
		{
			Filename: ".gitignore",
			Template: gitignore,
		},
		{
			Filename: "LICENSE",
			Template: licenseWithCopyright(),
		},
		{
			Filename: "README.md",
			Template: readme,
		},
		{
			Filename: "main.go",
			Template: main,
		},
		{
			Filename: "config/config.go",
			Template: config,
		},
		{
			Filename: "controllers/health.go",
			Template: health,
		},
		{
			Filename: "db/db.go",
			Template: db,
		},
		{
			Filename: "db/seed/development.yml",
			Template: "",
		},
		{
			Filename: "models/models.go",
			Template: models,
		},
	}
	if err := CreateAndAddFileList(files, owner, projectName); err != nil {
		for _, e := range err {
			fmt.Println(e)
		}
		panic("errors encountered while creating files")
	}

	runCommand("go", "mod", "init", "github.com/"+owner+"/"+projectName)
	runCommand("swag", "init")
	MustGitAddFile("docs/docs.go")
	MustGitAddFile("docs/swagger.json")
	MustGitAddFile("docs/swagger.yaml")
	runCommand("go", "mod", "tidy")
	MustGitAddFile("go.mod")
	MustGitAddFile("go.sum")
	runCommand("git", "commit", "-m", "\"Initial commit\"")
	runCommand("git", "flow", "init")
}
