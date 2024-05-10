package db

import (
	"bytes"
	"corefetch/identity/sys"
	"database/sql"
	"encoding/json"
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

func (a *Account) Display() sys.M {
	return sys.M{
		"id":    a.ID,
		"names": a.Names,
	}
}

func (a *Account) DisplayWithMeta(filter []string) sys.M {

	var matched = 0

	var meta = make(map[string]string)

	for key, value := range a.Meta {
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

func (a *Account) Save() error {

	sql := `
	INSERT INTO accounts
		(id, login, password, names, meta)
	VALUES
		($1, $2, $3, $4, $5)
	`

	buf := bytes.NewBufferString("")

	if err := json.NewEncoder(buf).Encode(a.Meta); err != nil {
		return err
	}

	_, err := db.Exec(
		sql,
		a.ID,
		a.Login,
		a.Password,
		strings.Join(a.Names, ","),
		buf.String(),
	)

	if err != nil {
		return err
	}

	return nil
}

func FindAccountByID(id int64) (account *Account, err error) {

	sql := `
	SELECT 
		id, login, password, names, meta, verified, created
	FROM
		accounts
	WHERE
		id = $1
	`

	return hydrate(db.QueryRow(sql, id))
}

func FindAccountByLogin(login string) (account *Account, err error) {

	sql := `
	SELECT 
		id, login, password, names, meta, verified, created
	FROM
		accounts
	WHERE
		login = $1
	`

	return hydrate(db.QueryRow(sql, login))
}

func hydrate(row *sql.Row) (account *Account, err error) {

	if row.Err() != nil {
		return nil, row.Err()
	}

	var names string
	var meta string
	var verified *time.Time

	account = &Account{}

	err = row.Scan(
		&account.ID,
		&account.Login,
		&account.Password,
		&names,
		&meta,
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

	if err := json.NewDecoder(bytes.NewBufferString(meta)).Decode(&account.Meta); err != nil {
		return nil, err
	}

	return
}
