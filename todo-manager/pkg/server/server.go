package server

import (
	"context"
	"errors"
	"fmt"

	todomgrpb "github.com/giantswarm/blog-i-want-it-all/todo-manager/pkg/proto"
	"github.com/jinzhu/gorm"

	// initialize mysql gorm driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const dbName = "todo"

var todo = &todomgrpb.Todo{
	Id:    100,
	Text:  "bogus",
	Done:  false,
	Owner: "test_user",
}

// TodoManagerServer implements gRPC server for todo manager
type TodoManagerServer struct {
	config *Config
	db     *gorm.DB
}

// NewTodoManagerServer creates a new TodoManagerServer
func NewTodoManagerServer(config *Config) *TodoManagerServer {
	mysqlConnectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MysqlUser, config.MysqlPass, config.MysqlHost, dbName)
	db, err := gorm.Open("mysql", mysqlConnectString)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}
	db.AutoMigrate(&TodoEntry{})

	mgr := &TodoManagerServer{
		config: config,
		db:     db,
	}
	return mgr
}

// Stop stops the server and frees resources
func (t *TodoManagerServer) Stop() {
	t.db.Close()
}

// CreateTodo stores new todo in database
func (t *TodoManagerServer) CreateTodo(ctx context.Context, todo *todomgrpb.Todo) (*todomgrpb.Todo, error) {
	dbTodo := FromGrpc(todo)
	t.db.Create(dbTodo)
	if dbTodo.ID == 0 {
		return nil, errors.New("Error inserting to database")
	}
	return dbTodo.ToGrpc(), nil
}

// ListTodos lists all todos owned by the user sent in request
func (t *TodoManagerServer) ListTodos(req *todomgrpb.ListTodosReq, srv todomgrpb.TodoManager_ListTodosServer) error {
	var todos []TodoEntry
	t.db.Where("owner = ?", req.Owner).Find(&todos)
	for _, t := range todos {
		todo := t.ToGrpc()
		srv.Send(todo)
	}
	return nil
}

// GetTodo returns todo with specified ID and owner, if it exists
func (t *TodoManagerServer) GetTodo(ctx context.Context, grpcTodo *todomgrpb.TodoIdReq) (*todomgrpb.Todo, error) {
	found := TodoEntry{}
	t.db.First(&found, grpcTodo.GetId())
	if found.ID == 0 || found.Owner != grpcTodo.GetOwner() {
		return nil, errors.New("Todo not found")
	}
	return found.ToGrpc(), nil
}

// UpdateTodo updates a todo with a specified ID and owner, if it exists
func (t *TodoManagerServer) UpdateTodo(ctx context.Context, grpcTodo *todomgrpb.Todo) (*todomgrpb.Todo, error) {
	found := TodoEntry{}
	t.db.First(&found, grpcTodo.GetId())
	if found.ID == 0 || found.Owner != grpcTodo.GetOwner() {
		return nil, errors.New("Todo not found")
	}

	found.Text = grpcTodo.Text
	found.Done = grpcTodo.Done
	t.db.Save(&found)
	if found.ID == 0 {
		return nil, errors.New("Error updating record in DB")
	}

	return found.ToGrpc(), nil
}

// DeleteTodo deletes a todo with a specified ID and owner, if it exists
func (t *TodoManagerServer) DeleteTodo(ctx context.Context, grpcTodo *todomgrpb.TodoIdReq) (*todomgrpb.DeleteTodoRes, error) {
	found := TodoEntry{}
	t.db.First(&found, grpcTodo.GetId())
	if found.ID == 0 || found.Owner != grpcTodo.GetOwner() {
		return nil, errors.New("Todo not found")
	}

	t.db.Delete(&found)

	return &todomgrpb.DeleteTodoRes{Success: true}, nil
}
