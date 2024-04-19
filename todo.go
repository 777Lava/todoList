package todo

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title"  db:"title" binding:"required"`
	Description string `json:"description"  db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int 
	ListId int
	ItemID int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done         *bool   `json:"done"`
}


func (i UpdateListInput) Validate() error{
	logrus.Print(i.Title,i.Description)
	if i.Title == nil && i.Description == nil {
		return errors.New("update structures has no values")
	}
	return nil 
}
func (i UpdateItemInput) Validate() error{
	logrus.Print(i.Title,i.Description)
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structures has no values")
	}
	return nil 
}
