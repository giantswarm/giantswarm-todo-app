package server

import (
	"context"
	"fmt"

	todomgrpb "github.com/giantswarm/blog-i-want-it-all/todo-manager/pkg/proto"
)

var todo = &todomgrpb.Todo{
	Id:    "100",
	Text:  "bogus",
	Done:  false,
	Owner: "test_user",
}

// TodoManagerServer implements gRPC server for todo manager
type TodoManagerServer struct{}

// CreateTodo stores new todo in database
func (*TodoManagerServer) CreateTodo(ctx context.Context, todo *todomgrpb.Todo) (*todomgrpb.Todo, error) {
	todo.Id = "101"
	return todo, nil
}

// ListTodos lists all todos owned by the user sent in request
func (*TodoManagerServer) ListTodos(req *todomgrpb.ListTodosReq, srv todomgrpb.TodoManager_ListTodosServer) error {
	srv.Send(todo)
	return nil
}

// GetTodo returns todo with specified ID and owner, if it exists
func (*TodoManagerServer) GetTodo(context.Context, *todomgrpb.TodoIdReq) (*todomgrpb.Todo, error) {
	return todo, nil
}

// UpdateTodo updates a todo with a specified ID and owner, if it exists
func (*TodoManagerServer) UpdateTodo(context.Context, *todomgrpb.Todo) (*todomgrpb.Todo, error) {
	todo.Text = fmt.Sprintf("%s-updated", todo.Text)
	return todo, nil
}

// DeleteTodo deletes a todo with a specified ID and owner, if it exists
func (*TodoManagerServer) DeleteTodo(context.Context, *todomgrpb.TodoIdReq) (*todomgrpb.DeleteTodoRes, error) {
	return &todomgrpb.DeleteTodoRes{}, nil
}
