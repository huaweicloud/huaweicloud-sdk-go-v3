package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateImageWatermarkRequest struct {
	Body *CreateImageWatermarkRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateImageWatermarkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkRequest struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkRequest", string(data)}, " ")
}
