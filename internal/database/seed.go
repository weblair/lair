package database

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"github.com/romanyx/polluter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"regexp"
	"strings"
)

func hashPassword(password []byte) []byte {
	r, err := regexp.Compile(`\$PASSWORD\((.*)\)`)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panic("Compiling regex failed.")
	}

	//noinspection GoNilness
	m := r.FindSubmatch(password)
	h, err := bcrypt.GenerateFromPassword(m[1], bcrypt.DefaultCost)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Creating password hash failed.")
	}

	logrus.WithFields(logrus.Fields{
		"password": string(password),
		"hash":     hex.EncodeToString(h),
	}).Debug("Password hashed.")

	return h
}

func processMacros(raw []byte) string {
	password, err := regexp.Compile(`\$PASSWORD\(.*\)`)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panic("Compiling regex failed.")
	}

	//noinspection GoNilness
	processed := password.ReplaceAllFunc(raw, hashPassword)

	return string(processed)
}

// SeedDatabase parses the seed YAML file for the given environment and populates the associated database.
func SeedDatabase(env string) {
	db, err := NewConnectionFromConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"error":       errors.WithStack(err),
		}).Fatal("Failed to create connection to database.")
	}

	logrus.WithFields(logrus.Fields{
		"environment": env,
	}).Info("Reading seed data for environment.")

	p := polluter.New(polluter.PostgresEngine(db))
	data := fmt.Sprintf("%s/%s.yml", viper.GetString("SEED_DIRECTORY"),  env)

	raw, err := ioutil.ReadFile(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"filename":    data,
			"error":       err,
		}).Fatal("Failed to open seed data file.")
	}
	processed := processMacros(raw)

	logrus.WithFields(logrus.Fields{
		"environment": env,
		"filename":    data,
	}).Info("Loading seed data into database.")
	if err := p.Pollute(strings.NewReader(processed)); err != nil {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"filename":    data,
			"error":       err,
		}).Fatal("Failed to load seed data into the database.")
	}
	logrus.WithFields(logrus.Fields{
		"environment": env,
		"filename":    data,
	}).Info("Database seeding complete.")
}
