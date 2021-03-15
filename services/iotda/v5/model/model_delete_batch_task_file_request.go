package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteBatchTaskFileRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	FileId     string  `json:"file_id"`
}

func (o DeleteBatchTaskFileRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBatchTaskFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteBatchTaskFileRequest", string(data)}, " ")
}
