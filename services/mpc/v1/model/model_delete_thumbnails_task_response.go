package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteThumbnailsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteThumbnailsTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteThumbnailsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteThumbnailsTaskResponse", string(data)}, " ")
}
