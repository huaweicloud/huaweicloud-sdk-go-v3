package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDocWatermarkRequest struct {
	Body *CreateDocWatermarkRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateDocWatermarkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDocWatermarkRequest struct{}"
	}

	return strings.Join([]string{"CreateDocWatermarkRequest", string(data)}, " ")
}
