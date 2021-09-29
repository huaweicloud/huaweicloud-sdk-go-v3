package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchResetServersPasswordRequest struct {
	Body *BatchResetServersPasswordRequestBody `json:"body,omitempty"`
}

func (o BatchResetServersPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchResetServersPasswordRequest struct{}"
	}

	return strings.Join([]string{"BatchResetServersPasswordRequest", string(data)}, " ")
}
