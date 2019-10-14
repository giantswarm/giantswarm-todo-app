package todo

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/piontec/go-chi-middleware-server/pkg/server/middleware"
	"google.golang.org/grpc"

	todomgrpb "github.com/giantswarm/blog-i-want-it-all/api-server/pkg/todo/proto"
)

// Username is a temporary value for all user name fields until we get proper authentication in place
const Username = "anonymous"

// Router is a registry of go-chi routes supported by Todo
type Router struct {
	grpcClient todomgrpb.TodoManagerClient
}

// NewRouter returns new go-chi router with initialized gRPC client
func NewRouter(todoManagerAddr string) *Router {
	requestOpts := grpc.WithInsecure()
	// Dial the server, returns a client connection
	conn, err := grpc.Dial(todoManagerAddr, requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to %s: %v", todoManagerAddr, err)
	}
	// Instantiate the TodoManagerClient with our client connection to the server
	client := todomgrpb.NewTodoManagerClient(conn)
	return &Router{
		grpcClient: client,
	}
}

// GetRouter returns configuredsub-router for Todo resources
func (t *Router) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", t.ListTodos)
	r.Post("/", t.CreateTodo) // POST /

	r.Route("/{todoID}", func(r chi.Router) {
		r.Get("/", t.GetTodo)       // GET /123
		r.Put("/", t.UpdateTodo)    // PUT /123
		r.Delete("/", t.DeleteTodo) // DELETE /123
	})

	return r
}

// ListTodos lists all todos owned by a user
func (t *Router) ListTodos(w http.ResponseWriter, r *http.Request) {
	stream, err := t.grpcClient.ListTodos(r.Context(), &todomgrpb.ListTodosReq{
		Owner: Username,
	})
	if err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return

	}
	for {
		res, err := stream.Recv()
		// If end of stream, break the loop
		if err == io.EOF {
			break
		}
		// if err, return an error
		if err != nil {
			render.Render(w, r, middleware.ErrRender(err))
			return
		}
		todo, _ := FromGRPCTodo(res)
		if err := render.Render(w, r, todo); err != nil {
			render.Render(w, r, middleware.ErrRender(err))
			return
		}
	}

}

// CreateTodo creates a new todo for a given user
func (t *Router) CreateTodo(w http.ResponseWriter, r *http.Request) {
	data := &Todo{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, &middleware.ErrResponse{HTTPStatusCode: 400, StatusText: "Invalid request."})
		return
	}
	fmt.Print(data.ID)
	fmt.Print(data.Text)
	fmt.Print(data.Done)
}

// GetTodo gets a todo with specified user and todo ID
func (t *Router) GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	fmt.Print(todoID)
}

// DeleteTodo deletes a todo with specified user and todo ID
func (t *Router) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	fmt.Print(todoID)
}

// UpdateTodo updates a todo with specified user and todo ID
func (t *Router) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	fmt.Print(todoID)
	data := &Todo{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, middleware.ErrInvalidRequest(err))
		return
	}
}
