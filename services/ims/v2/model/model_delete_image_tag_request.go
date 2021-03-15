package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteImageTagRequest struct {
	ImageId string `json:"image_id"`
	Key     string `json:"key"`
}

func (o DeleteImageTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteImageTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteImageTagRequest", string(data)}, " ")
}
