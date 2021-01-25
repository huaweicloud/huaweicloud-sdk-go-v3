package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateListenerRequest struct {
	Body *CreateListenerRequestBody `json:"body,omitempty"`
}

func (o CreateListenerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateListenerRequest struct{}"
	}

	return strings.Join([]string{"CreateListenerRequest", string(data)}, " ")
}
