package config

// TODO: Finish building this out for Gin project generation

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
)

func readCurrentDotEnv() (map[string]string, error) {
	values := make(map[string]string)

	f, err := os.Open(".env")
	if err != nil {
		// Failing to open the DotEnv file will result in assuming that a fresh one should be generated.
		return values, nil
	}
	//noinspection ALL
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.SplitN(s.Text(), "=", 2)
		if len(line) != 2 {
			return nil, errors.Errorf("invalid .env entry: %s", s.Text())
		}

		values[line[0]] = line[1]
	}

	return values, nil
}

// TODO: Copy config values from the config that match the DOTENV_KEYS list
func makeDotEnvMap(values *map[string]string) (*map[string]string, error) {
	return values, nil
}

func writeDotEnvFile(values map[string]string) error {
	contents := ""
	for k, v := range values {
		contents += fmt.Sprintf("%s=%s\n", k, v)
	}

	err := ioutil.WriteFile(".env", []byte(contents), 0600)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// CreateHMACKey creates a cryptographically secure random value for use by JWT generators.
func CreateHMACKey() (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// CreateGinDotEnv creates a basic .env file based on the current Viper settings.
// Values in an existing .env file are not touched.
func CreateGinDotEnv() {
	values, err := readCurrentDotEnv()
	if err != nil {
		panic(err)
	}
	key, err := CreateHMACKey()
	if err != nil {
		panic(err)
	}
	values["JWT_KEY"] = key

	err = writeDotEnvFile(values)
	if err != nil {
		panic(err)
	}
}
