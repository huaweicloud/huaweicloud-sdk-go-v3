package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePredefineTagsRequest struct {
	Body *ReqCreatePredefineTag `json:"body,omitempty"`
}

func (o CreatePredefineTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePredefineTagsRequest struct{}"
	}

	return strings.Join([]string{"CreatePredefineTagsRequest", string(data)}, " ")
}
