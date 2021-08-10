package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CompareFaceByFileRequest struct {
	Body *CompareFaceByFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CompareFaceByFileRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CompareFaceByFileRequest struct{}"
	}

	return strings.Join([]string{"CompareFaceByFileRequest", string(data)}, " ")
}
