package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMigprojectRequest struct {
	Body *PostMigProjectBody `json:"body,omitempty"`
}

func (o CreateMigprojectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMigprojectRequest struct{}"
	}

	return strings.Join([]string{"CreateMigprojectRequest", string(data)}, " ")
}
