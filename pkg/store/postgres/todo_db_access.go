package postgres

import "github.com/hakankaan/todo-api/pkg/store"

// GetTodo gets todo from db
func (s *service) GetTodo(guid string) (todo store.Todo, err error) {
	err = s.DB.Where("global = ?", guid).First(&todo).Error
	return
}

// AddTodo creates and add a new todo to db
func (s *service) AddTodo(t store.Todo) (err error) {

	err = s.DB.Create(&t).Error

	return
}

// MakeDone marks todo as done
func (s *service) MakeDone(guid string) (err error) {
	err = s.DB.Model(store.Todo{}).Where("global = ?", guid).UpdateColumn("is_done", true).Error
	return
}

// DeleteTodo removes todo from db
func (s *service) DeleteTodo(guid string) (err error) {
	// delete todo with global
	err = s.DB.Where("global = ?", guid).Delete(store.Todo{}).Error
	return
}
