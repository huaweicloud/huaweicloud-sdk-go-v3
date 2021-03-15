package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBatchTaskFilesRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
}

func (o ListBatchTaskFilesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBatchTaskFilesRequest struct{}"
	}

	return strings.Join([]string{"ListBatchTaskFilesRequest", string(data)}, " ")
}
