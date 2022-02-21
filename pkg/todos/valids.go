package todos

// addTodoRequest request type for addtodo route
type addTodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
