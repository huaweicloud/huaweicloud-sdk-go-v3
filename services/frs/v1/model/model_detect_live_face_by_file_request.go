package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectLiveFaceByFileRequest struct {
	Body *DetectLiveFaceByFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o DetectLiveFaceByFileRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByFileRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByFileRequest", string(data)}, " ")
}
