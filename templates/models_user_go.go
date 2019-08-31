package templates

// ModelsUserGo is the template for models/user.go which is created when the --auth flag is invoked.
const ModelsUserGo = `package models

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User represents a user of this API.
type User struct {
	ModelBase
	AcceptAfter time.Time ` + "`json:\"-\" gorm:\"not null\"`" + `
	Email string ` + "`json:\"email\" gorm:\"unique,not null\"`" + `
	EmailVerified bool ` + "`json:\"email_verified\" gorm:\"not null,default:false\"`" + `
	Password *string ` + "`json:\"password,omitempty\" gorm:\"-\"`" + `
	PasswordHash []byte ` + "`json:\"-\" gorm:\"column:password,not null\"`" + `
	PublicID uuid.UUID ` + "`json:\"public_id\" gorm:\"unique,not null\"`" + `
}

func (u *User) createPasswordHash() error {
	if u.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.Wrap(err, "creating password hash failed")
		}
		u.PasswordHash = hash
		u.Password = nil
		u.AcceptAfter = time.Now()
	}

	return nil
}

// BeforeSave executes before a user is created.
func (u *User) BeforeCreate() error {
	var verrs ValidationErrors

	err := u.createPasswordHash()
	if err != nil {
		return errors.WithStack(err)
	}

	if u.Password == nil {
		verrs = append(verrs, fmt.Sprintf("a password must be set to create a new user"))
	}

	return verrs
}

// BeforeSave executes before a user is updated.
func (u *User) BeforeSave() error {
	var verrs ValidationErrors

	err := u.createPasswordHash()
	if err != nil {
		return errors.WithStack(err)
	}

	if !govalidator.IsEmail(u.Email) {
		verrs = append(verrs, fmt.Sprintf("%s is not a valid email address", u.Email))
	}

	return verrs
}
`
