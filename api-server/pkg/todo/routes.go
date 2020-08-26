package todo

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/piontec/go-chi-middleware-server/pkg/server/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"

	// "go.opencensus.io/trace"
	"google.golang.org/grpc"

	todomgrpb "github.com/giantswarm/giantswarm-todo-app/api-server/pkg/todo/proto"
)

// Username is a temporary value for all user name fields until we get proper authentication in place
const Username = "anonymous"

var ()

// Router is a registry of go-chi routes supported by Todo
type Router struct {
	grpcClient       todomgrpb.TodoManagerClient
	getAllCounter    *prometheus.CounterVec
	getOneCounter    *prometheus.CounterVec
	deleteOneCounter *prometheus.CounterVec
	updateOneCounter *prometheus.CounterVec
	createOneCounter *prometheus.CounterVec
}

// NewRouter returns new go-chi router with initialized gRPC client
func NewRouter(todoManagerAddr string) *Router {
	// Dial the server, returns a client connection
	conn, err := grpc.Dial(todoManagerAddr, grpc.WithInsecure(), grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))
	if err != nil {
		log.Fatalf("Unable to establish client connection to %s: %v", todoManagerAddr, err)
	}
	// Instantiate the TodoManagerClient with our client connection to the server
	client := todomgrpb.NewTodoManagerClient(conn)
	return &Router{
		grpcClient: client,
		getAllCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Subsystem: "todo",
			Name:      "get_all_count_total",
			Help:      "The total number of successful GETs for all the todos of an user",
		}, []string{"user"}),
		getOneCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Subsystem: "todo",
			Name:      "get_one_count_total",
			Help:      "The total number of successful GETs for a single todo of an user",
		}, []string{"user"}),
		deleteOneCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Subsystem: "todo",
			Name:      "delete_one_count_total",
			Help:      "The total number of successful DELETEs for a single todo of an user",
		}, []string{"user"}),
		createOneCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Subsystem: "todo",
			Name:      "create_one_count_total",
			Help:      "The total number of successful POSTs for a single todo of an user",
		}, []string{"user"}),
		updateOneCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Subsystem: "todo",
			Name:      "update_one_count_total",
			Help:      "The total number of successful PUTs for a single todo of an user",
		}, []string{"user"}),
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
	var todoList []render.Renderer
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
		todoList = append(todoList, todo)
	}
	if err := render.RenderList(w, r, todoList); err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	t.getAllCounter.WithLabelValues(Username).Inc()
}

// CreateTodo creates a new todo for a given user
func (t *Router) CreateTodo(w http.ResponseWriter, r *http.Request) {
	// bind JSON from request to go object
	data := &Todo{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, middleware.ErrInvalidRequest(err))
		return
	}
	// validate - todo text can't be empty
	if data.Text == "" {
		render.Render(w, r, middleware.ErrInvalidRequest(errors.New("Text can't be empty")))
		return
	}
	// we don't have any real auth, let's pretend we always serve the user with ID 0
	data.ID = "0"
	// run request
	newGrpcTodo, err := t.grpcClient.CreateTodo(r.Context(), data.ToGRPCTodo(Username))
	if err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	// convert to JSON object and send response
	todo, _ := FromGRPCTodo(newGrpcTodo)
	if err := render.Render(w, r, todo); err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	t.createOneCounter.WithLabelValues(Username).Inc()
}

// GetTodo gets a todo with specified user and todo ID
func (t *Router) GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	id, err := strconv.ParseUint(todoID, 10, 64)
	if err != nil {
		render.Render(w, r, middleware.ErrInvalidRequest(err))
		return
	}
	grpcTodo, err := t.grpcClient.GetTodo(r.Context(), &todomgrpb.TodoIdReq{
		Id:    uint64(id),
		Owner: Username,
	})
	if err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	todo, _ := FromGRPCTodo(grpcTodo)
	if err := render.Render(w, r, todo); err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	t.getOneCounter.WithLabelValues(Username).Inc()
}

// DeleteTodo deletes a todo with specified user and todo ID
func (t *Router) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	id, err := strconv.ParseUint(todoID, 10, 64)
	if err != nil {
		render.Render(w, r, middleware.ErrInvalidRequest(err))
		return
	}
	deleteRes, err := t.grpcClient.DeleteTodo(r.Context(), &todomgrpb.TodoIdReq{
		Id:    uint64(id),
		Owner: Username,
	})
	if err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	if err := render.Render(w, r, FromGRPCDeleteRes(deleteRes)); err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	t.deleteOneCounter.WithLabelValues(Username).Inc()
}

// UpdateTodo updates a todo with specified user and todo ID
func (t *Router) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	_, err := strconv.ParseUint(todoID, 10, 64)
	if err != nil {
		render.Render(w, r, middleware.ErrInvalidRequest(err))
		return
	}
	data := &Todo{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, middleware.ErrInvalidRequest(err))
		return
	}
	data.ID = todoID
	if data.ID != "" && data.ID != todoID {
		render.Render(w, r, middleware.ErrInvalidRequest(errors.New("ID from JSON is not empty and doesn't match URL ID")))
		return
	}
	grpcTodo, err := t.grpcClient.UpdateTodo(r.Context(), data.ToGRPCTodo(Username))
	if err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	todo, _ := FromGRPCTodo(grpcTodo)
	if err := render.Render(w, r, todo); err != nil {
		render.Render(w, r, middleware.ErrRender(err))
		return
	}
	t.updateOneCounter.WithLabelValues(Username).Inc()
}
