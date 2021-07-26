package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateValueListRequest struct {
	Body *CreateValueListRequestBody `json:"body,omitempty"`
}

func (o CreateValueListRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateValueListRequest struct{}"
	}

	return strings.Join([]string{"CreateValueListRequest", string(data)}, " ")
}
