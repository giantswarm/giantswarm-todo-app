package todo

import (
	"net/http"
	todomgrpb "github.com/giantswarm/blog-i-want-it-all/api-server/pkg/todo/proto"
)

// Todo data model.
type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// Bind allows to set additional properties on Todo object; not used here
func (t *Todo) Bind(r *http.Request) error {
	return nil
}

// Render allows to modify the way Todo object is rendered to text; not used here
func (t *Todo) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ToGRPCTodo return gRPC DTO for the upstream todo-manager service
func (t *Todo) ToGRPCTodo(owner string) *todomgrpb.Todo {
	return &todomgrpb.Todo{
		Id: t.ID,
		Text: t.Text,
		Done: t.Done,
		Owner: owner,
	}
}

// FromGRPCTodo returns new Todo object and owner info based on gRPC DTO from the
// upstream todo-manager service
func FromGRPCTodo(grpcTodo *todomgrpb.Todo) (*Todo, string) {
	return &Todo{
		ID: grpcTodo.GetId(),
		Text: grpcTodo.GetText(),
		Done: grpcTodo.GetDone(),
	}, grpcTodo.GetOwner() 
}

// DeleteRes data model.
type DeleteRes struct {
	Success bool   `json:"success"`
}

// Bind allows to set additional properties on Todo object; not used here
func (t *DeleteRes) Bind(r *http.Request) error {
	return nil
}

// Render allows to modify the way Todo object is rendered to text; not used here
func (t *DeleteRes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// FromGRPCDeleteRes returns new DeleteRes object based on gRPC DTO from the
// upstream todo-manager service
func FromGRPCDeleteRes(grpcRes *todomgrpb.DeleteTodoRes) *DeleteRes {
	return &DeleteRes{
		Success: grpcRes.GetSuccess(),
	}
}