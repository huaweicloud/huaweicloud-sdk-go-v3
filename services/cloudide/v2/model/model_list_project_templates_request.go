package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectTemplatesRequest struct {
	Arch    *string `json:"arch,omitempty"`
	StackId string  `json:"stack_id"`
}

func (o ListProjectTemplatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTemplatesRequest", string(data)}, " ")
}
