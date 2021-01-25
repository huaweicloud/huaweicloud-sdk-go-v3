package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateResourceRequest struct {
	Body *ResourceInfo `json:"body,omitempty"`
}

func (o CreateResourceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceRequest", string(data)}, " ")
}
