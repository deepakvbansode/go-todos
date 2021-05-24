package models

type Todo struct {
	UserID    int         `json:"userId" bson:"userId"`
	ID        interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string      `json:"title" bson:"title"`
	Completed bool        `json:"completed" bson:"completed"`
}

type TodoUpdate struct {
	ModifiedCount int `json:"modifiedCount"`
	Result        Todo
}

type DeleteTodo struct {
	DeletedCount int `json:"deletedCount"`
}

type Session struct {
	UserID  string
	StarID  string
	EmailID string
}
