package redis

import (
	"context"
	"encoding/json"

	"github.com/hakankaan/todo-api/pkg/store"
)

var ctx = context.Background()

// AddTodo adds a todo to the database
func (s *service) AddTodo(todo store.Todo) (err error) {
	err = s.r.Set(ctx, todo.GuID, todo, 0).Err()
	return
}

// GetTodo gets a todo from the database
func (s *service) GetTodo(guid string) (todo store.Todo, err error) {
	val, err := s.r.Get(ctx, guid).Result()
	if err != nil {
		s.l.Error("s.r.Get", err)
		return
	}

	err = json.Unmarshal([]byte(val), &todo)

	return
}

// DeleteTodo deletes a todo from the database
func (s *service) DeleteTodo(guid string) (err error) {
	err = s.r.Del(ctx, guid).Err()
	return
}

// MakeDone marks a todo as done
func (s *service) MakeDone(guid string) (err error) {
	todo, err := s.GetTodo(guid)
	if err != nil {
		s.l.Error("s.GetTodo", err)
		return
	}
	todo.IsDone = true
	err = s.r.Set(ctx, guid, todo, 0).Err()
	return
}
