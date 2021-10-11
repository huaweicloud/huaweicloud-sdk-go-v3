package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEdgeNodeRequest struct {
	Body *EdgeNodeCreation `json:"body,omitempty"`
}

func (o CreateEdgeNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEdgeNodeRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeNodeRequest", string(data)}, " ")
}
