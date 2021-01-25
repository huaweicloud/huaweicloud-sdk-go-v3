package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchUpdateMembersRequest struct {
	Body *BatchUpdateMembersRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersRequest", string(data)}, " ")
}
