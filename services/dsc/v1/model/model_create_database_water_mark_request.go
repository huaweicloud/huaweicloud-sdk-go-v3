package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDatabaseWaterMarkRequest struct {
	Body *EmbeddedDatabaseWatermark `json:"body,omitempty"`
}

func (o CreateDatabaseWaterMarkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDatabaseWaterMarkRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseWaterMarkRequest", string(data)}, " ")
}
