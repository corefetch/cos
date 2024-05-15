package do

import (
	"fmt"
	"gom/api/messages"
	"gom/core/sys"
	"gom/pod/identity/db"
	"gom/pod/messages/ob"
	"net/mail"
	"os"
	"regexp"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(account *db.Account) (err error) {

	existent, err := db.FindAccountByLogin(account.Login)

	if err == nil || existent != nil {
		return errors.New("account with login already exists")
	}

	account.ID = sys.MustGenerateID()

	if _, err := mail.ParseAddress(account.Login); err != nil {
		return errors.Wrap(err, "invalid login email address")
	}

	// validate password and encrypt
	password, err := CreateSecurePassword(account.Password)

	if err != nil {
		return errors.Wrap(err, "invalid password")
	}

	// replace password in account
	account.Password = password

	if len(account.Names) < 2 {
		return errors.New("provide [first,last] names property")
	}

	if err := account.Save(); err != nil {
		return err
	}

	sys.Events().Publish("identity.created", []byte(fmt.Sprint(account.ID)))

	if os.Getenv("VERIFICATION") == "true" {

		sys.Logger().Info("Send verification code")

		err = messages.Send(account, "account.verify", &ob.Args{
			"CODE": "",
		})

		if err != nil {
			sys.Logger().Errorf("Failed to send verification code:", err.Error())
			return
		}
	}

	return
}

func CreateSecurePassword(password string) (pass string, err error) {

	const requirement = "password must contain at least a lowercase character, an uppercase character, and a digit"

	matchLower, _ := regexp.MatchString("[a-z]", password)
	matchUpper, _ := regexp.MatchString("[A-Z]", password)

	matchDigit, _ := regexp.MatchString("[0-9]", password)
	if !matchLower || !matchUpper || !matchDigit {
		return "", errors.New(requirement)
	}

	if len(password) < 8 {
		return "", errors.New("password must be bigger than 8 characters")
	}

	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bcryptHash), nil
}
