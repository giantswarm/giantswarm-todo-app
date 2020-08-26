package server

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	todomgrpb "github.com/giantswarm/giantswarm-todo-app/todo-manager/pkg/proto"
	"github.com/jinzhu/gorm"
	"go.opencensus.io/trace"

	// initialize mysql gorm driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const dbName = "todo"

// TodoManagerServer implements gRPC server for todo manager
type TodoManagerServer struct {
	config *Config
	db     *gorm.DB
}

// NewTodoManagerServer creates a new TodoManagerServer
func NewTodoManagerServer(config *Config) *TodoManagerServer {
	driverName := "mysql"
	mysqlConnectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MysqlUser, config.MysqlPass, config.MysqlHost, dbName)
	db, err := gorm.Open(driverName, mysqlConnectString)
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
	_, span := trace.StartSpan(ctx, "db-create")
	t.db.Create(dbTodo)
	span.End()
	if dbTodo.ID == 0 {
		return nil, errors.New("Error inserting to database")
	}
	return dbTodo.ToGrpc(), nil
}

// ListTodos lists all todos owned by the user sent in request
func (t *TodoManagerServer) ListTodos(req *todomgrpb.ListTodosReq, srv todomgrpb.TodoManager_ListTodosServer) error {
	var todos []TodoEntry
	_, span := trace.StartSpan(srv.Context(), "db-list")
	if t.config.EnableFailures {
		if num := rand.Int() % 10; num == 0 {
			// simulate the DB is really slow
			time.Sleep(time.Duration(rand.Int()%3+1) * time.Second)
		}
	}
	t.db.Where("owner = ?", req.Owner).Find(&todos)
	span.End()
	for _, t := range todos {
		todo := t.ToGrpc()
		srv.Send(todo)
	}
	return nil
}

// GetTodo returns todo with specified ID and owner, if it exists
func (t *TodoManagerServer) GetTodo(ctx context.Context, grpcTodo *todomgrpb.TodoIdReq) (*todomgrpb.Todo, error) {
	found := TodoEntry{}
	_, span := trace.StartSpan(ctx, "db-get")
	t.db.First(&found, grpcTodo.GetId())
	span.End()
	if found.ID == 0 || found.Owner != grpcTodo.GetOwner() {
		return nil, errors.New("Todo not found")
	}
	return found.ToGrpc(), nil
}

// UpdateTodo updates a todo with a specified ID and owner, if it exists
func (t *TodoManagerServer) UpdateTodo(ctx context.Context, grpcTodo *todomgrpb.Todo) (*todomgrpb.Todo, error) {
	found := TodoEntry{}
	_, span := trace.StartSpan(ctx, "db-update-get")
	t.db.First(&found, grpcTodo.GetId())
	span.End()
	if found.ID == 0 || found.Owner != grpcTodo.GetOwner() {
		return nil, errors.New("Todo not found")
	}

	found.Text = grpcTodo.Text
	found.Done = grpcTodo.Done
	_, span = trace.StartSpan(ctx, "db-update-save")
	t.db.Save(&found)
	span.End()
	if found.ID == 0 {
		return nil, errors.New("Error updating record in DB")
	}

	return found.ToGrpc(), nil
}

// DeleteTodo deletes a todo with a specified ID and owner, if it exists
func (t *TodoManagerServer) DeleteTodo(ctx context.Context, grpcTodo *todomgrpb.TodoIdReq) (*todomgrpb.DeleteTodoRes, error) {
	found := TodoEntry{}
	_, span := trace.StartSpan(ctx, "db-update-get")
	t.db.First(&found, grpcTodo.GetId())
	span.End()
	if found.ID == 0 || found.Owner != grpcTodo.GetOwner() {
		return nil, errors.New("Todo not found")
	}

	_, span = trace.StartSpan(ctx, "db-update-delete")
	t.db.Delete(&found)
	span.End()

	return &todomgrpb.DeleteTodoRes{Success: true}, nil
}
