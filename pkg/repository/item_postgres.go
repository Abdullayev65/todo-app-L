package repository

import (
	"errors"
	"fmt"
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Create(userId int, listId int, list todo.TodoItem) (int, error) {
	// check list owner
	{
		query := fmt.Sprintf(`SELECT l.id FROM %s l INNER JOIN %s ul ON ul.list_id = l.id
 INNER JOIN %s u ON ul.user_id = u.id WHERE u.id = $1 AND l.id=$2`,
			todoListsTable, usersTodoListsTable, usersTable)
		row := r.db.QueryRow(query, userId, listId)
		var l_id int
		err := row.Scan(&l_id)
		if err != nil {
			// TODO userga tegishli bomaganligi haqida aytish
			return 0, err
		}
		if l_id != listId {
			return 0, errors.New("something wrong")
		}
	}
	// create item
	{
		tx, err := r.db.Begin()
		if err != nil {
			return 0, err
		}

		query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemTable)
		row := tx.QueryRow(query, list.Title, list.Description)
		var itemId int
		if err = row.Scan(&itemId); err != nil {
			tx.Rollback()
			return 0, err
		}
		usersListQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
		_, err = tx.Exec(usersListQuery, listId, itemId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		return itemId, tx.Commit()
	}
}

func (r *ItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	query := fmt.Sprintf(`SELECT i.* FROM %s i INNER JOIN %s il ON il.item_id=i.id
 INNER JOIN %s l ON il.list_id=l.id INNER JOIN %s ul ON ul.list_id=l.id WHERE user_id = $1 AND l.id = $2`,
		todoItemTable, listsItemsTable, todoListsTable, usersTodoListsTable)
	var lists []todo.TodoItem
	err := r.db.Select(&lists, query, userId, listId)
	return lists, err
}

func (r *ItemPostgres) GetById(userId, listId, itemId int) (todo.TodoItem, error) {
	query := fmt.Sprintf(`SELECT i.* FROM %s i INNER JOIN %s il ON il.item_id=i.id
 INNER JOIN %s l ON il.list_id=l.id INNER JOIN %s ul ON ul.list_id=l.id 
 WHERE user_id = $1 AND l.id = $2 AND i.id = $3`,
		todoItemTable, listsItemsTable, todoListsTable, usersTodoListsTable)
	var item todo.TodoItem
	err := r.db.Get(&item, query, userId, listId, itemId)
	return item, err
}
func (r *ItemPostgres) Delete(userId, listId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemTable, listsItemsTable, usersTodoListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	fmt.Println(query)
	return err
}
func (r *ItemPostgres) Update(userId int, listId int, itemId int, input todo.InputItemUpdate) error {
	argNum, args, setValues := 1, *new([]any), *new([]string)
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argNum))
		args = append(args, *input.Title)
		argNum++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argNum))
		args = append(args, *input.Description)
		argNum++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argNum))
		args = append(args, *input.Done)
		argNum++
	}
	if argNum == 1 {
		return errors.New("all fields are nil")
	}
	setStr := strings.Join(setValues, " ,")
	query := fmt.Sprintf(`UPDATE %s i SET %s FROM %s il 
INNER JOIN %s l ON il.list_id=l.id INNER JOIN %s ul ON ul.list_id=l.id 
 WHERE user_id = $%d AND l.id = $%d AND i.id = $%d AND il.item_id=i.id`,
		todoItemTable, setStr, listsItemsTable,
		todoListsTable, usersTodoListsTable, argNum, argNum+1, argNum+2)
	args = append(args, userId, listId, itemId)
	_, err := r.db.Exec(query, args...)
	return err
}
