package data

type Todo struct {
	ID   string
	Item string
	Done bool
}

type UserTodos struct {
	SessionID string
	Todos     []Todo
}
