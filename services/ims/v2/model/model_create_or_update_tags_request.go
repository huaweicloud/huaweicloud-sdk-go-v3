package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateOrUpdateTagsRequest struct {
	Body *AddOrUpdateTagsRequestBody `json:"body,omitempty"`
}

func (o CreateOrUpdateTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateOrUpdateTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateTagsRequest", string(data)}, " ")
}
