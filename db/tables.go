package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var (
	postgresUser     = os.Getenv("POSTGRES_USER")
	postgresHost     = os.Getenv("POSTGRES_HOST")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase = os.Getenv("POSTGRES_DATABASE")
	//
	postgresURL = fmt.Sprintf("postgres://%s:%s@%s/%s", postgresUser,
		postgresPassword, postgresHost, postgresDatabase)
)

// Db
type Db struct {
	conn *pgx.Conn
}

// NewDB
func NewDB() (Db, error) {
	conn, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		return Db{}, err
	}

	return Db{
		conn: conn,
	}, nil
}

// Close
func (d Db) Close() error {
	return d.conn.Close(context.TODO())
}

// InitialDatabase
func (d Db) InitialDatabase() error {
	_, err := d.conn.Exec(context.TODO(), `
CREATE TABLE IF NOT EXISTS groups (
	id varchar primary key,
	name varchar not null,
	user_ids varchar -- user ids separated by comma.
);

CREATE TABLE IF NOT EXISTS users (
	id varchar primary key,
	username varchar not null,
	password varchar not null,
	email varchar
);
`)
	if err != nil {
		return err
	}

	return nil
}
