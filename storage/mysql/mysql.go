package mysql

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jccatrinck/cartesian/libs/env"
	"github.com/jccatrinck/cartesian/storage/mysql/statements"
	"github.com/jmoiron/sqlx"
)

const (
	errDuplicate = 1062
)

// MySQL implements storage.Storage interface
type MySQL struct {
	db *sqlx.DB
}

// New creates a new instance of MySQL
func New() (m *MySQL, err error) {
	root, err := connect("root", env.Get("MYSQL_ROOT_PASSWORD", ""))

	if err != nil {
		return
	}

	err = exec(root, statements.Schema)

	if err != nil {
		return
	}

	m = &MySQL{}

	m.db, err = connect(
		env.Get("MYSQL_USER", "root"),
		env.Get("MYSQL_PASSWORD", ""),
	)

	if err != nil {
		return
	}

	return
}

func connect(user, password string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&timeout=120s",
		user,
		password,
		env.Get("MYSQL_HOST", "localhost"),
		env.Get("MYSQL_PORT", "3306"),
		env.Get("MYSQL_DATABASE", "api"),
	)

	return sqlx.Connect("mysql", dsn)
}

func exec(db *sqlx.DB, commands string) (err error) {
	for _, cmd := range strings.Split(commands, ";") {
		cmd = strings.TrimSpace(cmd)

		if cmd == "" {
			continue
		}

		_, err = db.Exec(cmd)

		if err != nil {
			return
		}
	}

	return
}
