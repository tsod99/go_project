package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

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

var (
	ErrConflict  = errors.New("Error conflict")
	ErrInsertion = errors.New("Error insert new element into the database")
	ErrUpdate    = errors.New("Error failed to update element in the table")
	ErrDelete    = errors.New("Error failed to delete element from the table")
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
	name varchar not null
);

CREATE TABLE IF NOT EXISTS users (
	id varchar primary key,
	username varchar not null,
	password varchar not null,
	email varchar default '',
	group_id varchar default ''
);
`)
	if err != nil {
		return err
	}

	return nil
}

// Group
type Group struct {
	Id   string
	Name string
}

// ListGroups
func (d Db) ListGroups() ([]Group, error) {
	rows, err := d.conn.Query(context.TODO(), `SELECT id, name FROM groups`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := []Group{}

	for rows.Next() {
		var group Group
		if err := rows.Scan(&group.Id, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func GetGroupByID(id string) (g Group, err error) {
	db, err := NewDB()
	if err != nil {
		return Group{}, err
	}
	defer db.Close()
	err = db.conn.QueryRow(context.TODO(), `
SELECT id, name FROM groups
WHERE id = $1`, id).Scan(&g.Id, &g.Name)
	if err != nil && err != pgx.ErrNoRows {
		return
	}
	err = nil
	return
}

// GetGroupIDByName
func GetGroupIDByName(name string) (g Group, err error) {
	db, err := NewDB()
	if err != nil {
		return Group{}, err
	}
	defer db.Close()
	err = db.conn.QueryRow(context.TODO(), `
SELECT id, name FROM groups
WHERE name = $1`, name).Scan(&g.Id, &g.Name)
	if err != nil && err != pgx.ErrNoRows {
		return
	}
	err = nil
	return
}

type User struct {
	Id       string
	Username string
	Password string
	Email    string
	Group    string
}

// ListUsers
func (d Db) ListUsers() ([]User, error) {
	rows, err := d.conn.Query(context.TODO(), `
SELECT id, username, password, email, group_id
FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		var id string
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Email,
			&id,
		)
		if err != nil {
			return nil, err
		}
		id = strings.TrimSpace(id)
		// fmt.Println("debug", id)
		g, err := GetGroupByID(id)
		if err != nil {
			return nil, err
		}
		// fmt.Println("debug", g)
		user.Group = g.Name
		// fmt.Println("debug", user)
		users = append(users, user)
	}

	return users, nil
}

// GetUser
func GetUser(username string) (User, error) {
	db, err := NewDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	var user User
	if err := db.conn.QueryRow(context.TODO(), `
SELECT id, username, password, email, group_id
FROM users WHERE username = $1`, username).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Group,
	); err != nil && err != pgx.ErrNoRows {
		return User{}, err
	}

	return user, nil
}

// CreateUser
func (d Db) CreateUser(u User) error {
	testUser, err := GetUser(u.Username)
	if err != nil {
		return err
	}
	if testUser.Username == u.Username {
		return ErrConflict
	}
	g, err := GetGroupIDByName(u.Group)
	if err != nil {
		return err
	}
	commandTag, err := d.conn.Exec(context.TODO(), `
INSERT INTO users(id, username, password, email, group_id)
VALUES($1, $2, $3, $4, $5)`, u.Id, u.Username, u.Password, u.Email, g.Id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrInsertion
	}
	return nil
}

// CreateGroup
func (d Db) CreateGroup(g Group) error {
	testGroup, err := GetGroupIDByName(g.Name)
	if err != nil {
		return err
	}
	if testGroup.Name == g.Name {
		return ErrConflict
	}
	commandTag, err := d.conn.Exec(context.TODO(), `
INSERT INTO groups(id, name)
VALUES($1, $2)`, g.Id, g.Name)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return ErrInsertion
	}
	return nil
}

// UpdateUser
func (d Db) UpdateUser(username, field, value string) error {
	var query string
	if field == "group" {
		field = "group_id"
		g, err := GetGroupIDByName(value)
		if err != nil {
			return err
		}
		value = g.Id
		query = fmt.Sprintf("UPDATE users set %s = $1 WHERE username = $2 ", field)
	} else {
		query = fmt.Sprintf("UPDATE users set %s = $1 WHERE username = $2 ", field)
	}
	commandTag, err := d.conn.Exec(context.TODO(), query, value, username)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrUpdate
	}
	return nil
}

// UpdateGroup
func (d Db) UpdateGroup(groupName, newGroupName string) error {
	commandTag, err := d.conn.Exec(context.TODO(),
		`UPDATE groups SET name = $1 where name = $2`,
		newGroupName, groupName)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrUpdate
	}
	return nil
}

// DeleteUser
func (d Db) DeleteUser(username string) error {
	commandTag, err := d.conn.Exec(context.TODO(), `DELETE FROM users WHERE username = $1`, username)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrDelete
	}
	return nil
}

// DeleteGroup
func (d Db) DeleteGroup(groupName string) error {
	commandTag, err := d.conn.Exec(context.TODO(), `DELETE FROM groups WHERE name = $1`, groupName)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrDelete
	}
	return nil
}
