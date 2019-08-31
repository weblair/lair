package generator

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"

	"github.com/weblair/lair/db"
	"github.com/weblair/lair/fileutil"
	"github.com/weblair/lair/templates"
)

type logStream struct {
	fields logrus.Fields
	level  logrus.Level
}

func (l *logStream) Write(p []byte) (int, error) {
	switch l.level {
	case logrus.InfoLevel:
		logrus.WithFields(l.fields).Info(string(p))
	case logrus.ErrorLevel:
		logrus.WithFields(l.fields).Error(string(p))
	default:
		return 0, errors.New("logStream initialized with an unsupported loglevel")
	}

	return len(p), nil
}

func runCommand(command string, args ...string) {
	fields := logrus.Fields{
		"command": command,
		"args":    args,
	}
	out := &logStream{fields: fields, level: logrus.InfoLevel}

	cmd := exec.Command(command, args...)
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		logrus.WithFields(logrus.Fields{
			"command": command,
			"error": err,
		}).Panic("Failed to execute command.")
	}
}

func setupGitFlow(path string) error {
	f, err := os.OpenFile(path+"/.git/config", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.WithMessage(err, "unable to open Git config")
	}

	if _, err := f.WriteString(templates.GitFlowGitconfig); err != nil {
		return errors.WithMessage(err, "could not append Git Flow config to Git config file")
	}

	return nil
}

// NewGinProject creates a runnable Gin project with database.
func NewGinProject(projectName string, withAuth bool) {
	owner := viper.GetString("GITHUB_USER")
	if owner == "" {
		logrus.Fatal("GITHUB_USER is not set in root config.")
	}

	u, err := user.Current()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Unable to get current user.")
	}

	//noinspection GoNilness
	projDir := u.HomeDir + "/devel/prog/" + projectName
	logrus.WithFields(logrus.Fields{
		"directory": projDir,
	}).Info("Creating project directory.")
	if err := os.Mkdir(projDir, 0755); err != nil {
		logrus.WithFields(logrus.Fields{
			"directory": projDir,
			"error":     errors.WithStack(err),
		}).Fatal("Failed to create project directory.")
	}
	if err := os.Chdir(projDir); err != nil {
		logrus.WithFields(logrus.Fields{
			"directory": projDir,
			"error":     errors.WithStack(err),
		}).Fatal("Failed to enter project directory.")
	}

	// TODO: Run Git commands through a Go library instead of a system call
	logrus.WithFields(logrus.Fields{
		"directory": projDir,
	}).Info("Initializing Git repository.")
	runCommand("git", "init")
	if err := setupGitFlow(projDir); err != nil {
		logrus.WithFields(logrus.Fields{
			"directory": projDir,
			"error":     errors.WithStack(err),
		}).Fatal("Failed to initialize Git repository.")
	}

	dirs := []fileutil.DirectoryData{
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
			Empty:   !withAuth,
		},
		{
			Dirname: "db/seed",
			Empty:   false,
		},
		{
			Dirname: "middleware",
			Empty:   !withAuth,
		},
		{
			Dirname: "models",
			Empty:   false,
		},
		{
			Dirname: "tests",
			Empty:   true,
		},
	}
	if err := fileutil.CreateAndAddDirectoryList(dirs); err != nil {
		logrus.Fatal("Errors encountered while creating directories.")
	}

	files := []fileutil.FileData{
		{
			Filename: ".gitignore",
			Template: templates.Gitignore,
		},
		{
			Filename: "LICENSE",
			Template: templates.LicenseWithCopyright(),
		},
		{
			Filename: "README.md",
			Template: templates.ReadmeMd,
		},
		{
			Filename: "Makefile",
			Template: templates.Makefile,
		},
		{
			Filename: ".air.conf",
			Template: templates.Dotair,
		},
		{
			Filename: "main.go",
			Template: templates.MainGo,
		},
		{
			Filename: "config/config.go",
			Template: templates.ConfigGo,
		},
		{
			Filename: "config/development.yml",
			Template: templates.ConfigDevelopmentYml,
		},
		{
			Filename: "controllers/health.go",
			Template: templates.HealthGo,
		},
		{
			Filename: "db/db.go",
			Template: templates.DbGo,
		},
		{
			Filename: "db/seed/development.yml",
			Template: "",
		},
		{
			Filename: "models/models.go",
			Template: templates.ModelsGo,
		},
	}

	if withAuth {
		authFiles := []fileutil.FileData{
			{
				Filename: "db/migrations/000001_create_users.up.sql",
				Template: templates.CreateUsersUp,
			},
			{
				Filename: "db/migrations/000001_create_users.down.sql",
				Template: templates.CreateUsersDown,
			},
			{
				Filename: "models/user.go",
				Template: templates.ModelsUserGo,
			},
		}

		files = append(files, authFiles...)
	}

	if err := fileutil.CreateAndAddFileList(files, owner, projectName); err != nil {
		logrus.Fatal("Errors encountered while creating files.")
	}
	// The .env file should not be added to Git, so we do one extra file create call here.
	logrus.WithFields(logrus.Fields{
		"file": ".env",
	}).Info("Creating file.")
	if err := ioutil.WriteFile(".env", []byte(templates.Dotenv), 0644); err != nil {
		logrus.WithFields(logrus.Fields{
			"file":  ".env",
			"error": err,
		}).Info("Failed to create file.")
	}

	runCommand("go", "mod", "init", "github.com/"+owner+"/"+projectName)
	runCommand("swag", "init")
	fileutil.MustGitAddFile("docs/docs.go")
	fileutil.MustGitAddFile("docs/swagger.json")
	fileutil.MustGitAddFile("docs/swagger.yaml")
	runCommand("go", "mod", "tidy")
	fileutil.MustGitAddFile("go.mod")
	fileutil.MustGitAddFile("go.sum")
	runCommand("git", "commit", "-m", "\"Initial commit\"")
	runCommand("git", "checkout", "-b", "develop")

	if err := db.CreateDatabaseWithName(projectName+"_development", false); err != nil {
		panic(errors.WithMessage(err, "failed to create dev database"))
	}
}
