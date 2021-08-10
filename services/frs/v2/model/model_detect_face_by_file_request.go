package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectFaceByFileRequest struct {
	Body *DetectFaceByFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o DetectFaceByFileRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectFaceByFileRequest struct{}"
	}

	return strings.Join([]string{"DetectFaceByFileRequest", string(data)}, " ")
}
