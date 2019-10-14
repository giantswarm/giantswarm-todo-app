package todo

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/piontec/go-chi-middleware-server/pkg/server/middleware"
)

type Router struct {
	// grpcClient clnmgrpb.ClientManagerClient
}

func NewRouter(todoManagerAddr string) *Router {
	// requestOpts := grpc.WithInsecure()
	// // Dial the server, returns a client connection
	// conn, err := grpc.Dial(clientManagerAddr, requestOpts)
	// if err != nil {
	//     log.Fatalf("Unable to establish client connection to %s: %v", clientManagerAddr, err)
	// }
	// // Instantiate the BlogServiceClient with our client connection to the server
	// client := clnmgrpb.NewClientManagerClient(conn)
	return &Router{
		// grpcClient: client,
	}
}

func (t *Router) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", t.ListTodos)
	r.Post("/", t.CreateTodo) // POST /articles

	r.Route("/{todoID}", func(r chi.Router) {
		r.Get("/", t.GetTodo)       // GET /articles/123
		r.Put("/", t.UpdateTodo)    // PUT /articles/123
		r.Delete("/", t.DeleteTodo) // DELETE /articles/123
	})

	return r
}

func (t *Router) ListTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Print("create")
}

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

func (t *Router) GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	fmt.Print(todoID)
}

func (t *Router) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	fmt.Print(todoID)
}

func (t *Router) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	fmt.Print(todoID)
	data := &Todo{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, &middleware.ErrResponse{HTTPStatusCode: 400, StatusText: "Invalid request."})
		return
	}
}
