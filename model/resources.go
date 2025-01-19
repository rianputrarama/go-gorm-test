package model

import (
	"time"
)

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Error     string      `json:"error,omitempty"`
	Token     string      `json:"token,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp string      `json:"timestamp"`
}

func (r *Response) SetResponse(s int, msg, err string) {
	r.Status = s
	r.Message = msg
	r.Error = err
	r.Timestamp = time.Now().Format("2006-01-02 15:04:05")
}
