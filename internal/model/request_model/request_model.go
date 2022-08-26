package request_model

type RequestId struct {
	Id string
}

type RequestTodo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
