package do

import (
	"edx/pod/identity/db"
	"edx/pod/identity/sys"
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
	password, err := createSecurePassword(account.Password)

	if err != nil {
		return errors.Wrap(err, "invalid password")
	}

	if len(account.Names) < 2 {
		return errors.New("provide [first,last] names property")
	}

	// replace password in account
	account.Password = password

	if err := account.Save(); err != nil {
		return err
	}

	if os.Getenv("VERIFICATION") == "true" {
		sys.Logger().Info("Send verification code")
	}

	return
}

func createSecurePassword(password string) (pass string, err error) {

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
