package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddImageTagRequest struct {
	ImageId string                  `json:"image_id"`
	Body    *AddImageTagRequestBody `json:"body,omitempty"`
}

func (o AddImageTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddImageTagRequest struct{}"
	}

	return strings.Join([]string{"AddImageTagRequest", string(data)}, " ")
}
