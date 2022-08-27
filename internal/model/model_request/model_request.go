package model_request

type RequestId struct {
	Id string
}

type RequestTodo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
