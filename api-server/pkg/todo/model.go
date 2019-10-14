package todo

import "net/http"

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
