package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteMembersRequest struct {
	Body *BatchAddMembersRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersRequest", string(data)}, " ")
}
