package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/avijit-dhaliwal/RESTfulAPI/internal/models"
)

type Database struct {
	*sql.DB
}

func InitDB() (*Database, error) {
	connStr := "user=username dbname=dbname sslmode=disable password=password"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (db *Database) GetItems() ([]models.Item, error) {
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var i models.Item
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func (db *Database) CreateItem(item *models.Item) error {
	err := db.QueryRow("INSERT INTO items(name) VALUES($1) RETURNING id", item.Name).Scan(&item.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetItem(id int) (models.Item, error) {
	var i models.Item
	err := db.QueryRow("SELECT id, name FROM items WHERE id = $1", id).Scan(&i.ID, &i.Name)
	if err != nil {
		return models.Item{}, err
	}
	return i, nil
}

func (db *Database) UpdateItem(item *models.Item) error {
	_, err := db.Exec("UPDATE items SET name = $1 WHERE id = $2", item.Name, item.ID)
	return err
}

func (db *Database) DeleteItem(id int) error {
	_, err := db.Exec("DELETE FROM items WHERE id = $1", id)
	return err
}