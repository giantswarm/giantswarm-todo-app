package server

import (
	"github.com/jinzhu/gorm"

	todomgrpb "github.com/giantswarm/blog-i-want-it-all/todo-manager/pkg/proto"
)

// TodoEntry is an object used for ORM mapping into the DB
type TodoEntry struct {
	gorm.Model
	Text  string
	Done  bool
	Owner string
}

// ToGrpc returns GRPC object from DB object
func (e *TodoEntry) ToGrpc() *todomgrpb.Todo {
	return &todomgrpb.Todo{
		Id:    uint64(e.ID),
		Text:  e.Text,
		Done:  e.Done,
		Owner: e.Owner,
	}
}

// FromGrpc returns DB object from GRPC object
func FromGrpc(grpcTodo *todomgrpb.Todo) *TodoEntry {
	return &TodoEntry{
		Model: gorm.Model{ID: uint(grpcTodo.GetId())},
		Text:  grpcTodo.GetText(),
		Done:  grpcTodo.GetDone(),
		Owner: grpcTodo.GetOwner(),
	}
}
