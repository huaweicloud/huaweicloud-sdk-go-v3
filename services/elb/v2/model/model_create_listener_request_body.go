package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateListenerRequestBody struct {
	Listener *CreateListenerReq `json:"listener"`
}

func (o CreateListenerRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateListenerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateListenerRequestBody", string(data)}, " ")
}
