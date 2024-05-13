package db

import (
	"database/sql"
	"edx/core/sys"
	"fmt"
	"strings"
	"time"
)

type Account struct {
	ID       int64             `json:"id"`
	Login    string            `json:"login"`
	Password string            `json:"password"`
	Names    []string          `json:"names"`
	Meta     map[string]string `json:"meta"`
	Verified *time.Time        `json:"verified,omitempty" bson:"verified,omitempty"`
	Created  time.Time         `json:"created"`
}

func (a *Account) SendTo() string {
	return a.Login
}

func (a *Account) SendName() []string {
	return a.Names
}

func (a *Account) Display() sys.M {
	return sys.M{
		"id":    a.ID,
		"names": a.Names,
	}
}

func (a *Account) DisplayWithMeta(filter []string) sys.M {

	var matched = 0

	var meta = make(map[string]string)

	for key, value := range a.GetMetas() {
		for _, filterKey := range filter {
			if key == filterKey {
				meta[filterKey] = value
				matched++
			}
		}
	}

	out := sys.M{
		"id":    a.ID,
		"names": a.Names,
	}

	if matched > 0 {
		out["meta"] = meta
	}

	return out
}

func (a *Account) GetMetas() (metas map[string]string) {

	metas = make(map[string]string)

	rows, err := db.Query("SELECT name, value FROM meta WHERE account=$1", a.ID)

	if err != nil {
		sys.Logger().Error("failed to query metas: ", err.Error())
		return metas
	}

	for rows.Next() {

		var k, v string

		if err := rows.Scan(&k, &v); err != nil {
			sys.Logger().Error("failed to get metas: ", err.Error())
			return metas
		}

		metas[k] = v
	}

	return
}

func (a *Account) GetMeta(name string) (v string, exist bool) {

	res := db.QueryRow("SELECT value FROM meta WHERE account=$1 and name=$2", a.ID, name)

	if res.Err() != nil {
		return v, false
	}

	if err := res.Scan(&v); err != nil {
		return "", false
	}

	return v, true
}

func (a *Account) updateMeta() {

	for k, v := range a.Meta {

		if _, exist := a.GetMeta(k); exist {

			sql := `
				UPDATE meta 
					SET value = $1
				WHERE
					account = $2
				AND
					name = $3
			`

			if _, err := db.Exec(sql, v, a.ID, k); err != nil {
				sys.Logger().Errorf("failed to update meta: %s", err.Error())
			}

		} else {

			sql := `
				INSERT INTO meta (account,name,value)
				VALUES (` + fmt.Sprintf("%d, '%s','%s'", a.ID, k, v) + `)
			`

			if _, err := db.Exec(sql); err != nil {
				sys.Logger().Errorf("failed to add meta: %s", err.Error())
			}
		}

	}
}

func (a *Account) Save() error {

	var sql string

	if existent, _ := FindAccountByID(a.ID); existent != nil {

		sql = `
		UPDATE
			accounts
		SET
			names = $1,
			password = $2
		WHERE
			id = $3
		`

		_, err := db.Exec(
			sql,
			strings.Join(a.Names, ","),
			a.Password,
			a.ID,
		)

		a.updateMeta()

		return err
	}

	sql = `
	INSERT INTO accounts
		(id, login, password, names)
	VALUES
		($1, $2, $3, $4)
	`

	_, err := db.Exec(
		sql,
		a.ID,
		a.Login,
		a.Password,
		strings.Join(a.Names, ","),
	)

	if err != nil {
		return err
	}

	a.updateMeta()

	return nil
}

func FindAccountByID(id int64) (account *Account, err error) {
	return hydrate(db.QueryRow(`SELECT * FROM accounts WHERE id = $1`, id))
}

func FindAccountByLogin(login string) (account *Account, err error) {
	return hydrate(db.QueryRow(`SELECT * FROM accounts WHERE login = $1`, login))
}

func hydrate(row *sql.Row) (account *Account, err error) {

	if row.Err() != nil {
		return nil, row.Err()
	}

	var names string
	var verified *time.Time

	account = &Account{}

	err = row.Scan(
		&account.ID,
		&account.Login,
		&account.Password,
		&names,
		&verified,
		&account.Created,
	)

	if err != nil {
		return nil, err
	}

	if verified != nil && !verified.IsZero() {
		account.Verified = verified
	}

	account.Names = strings.Split(names, ",")

	return
}
