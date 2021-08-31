package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectLiveByFileRequest struct {
	Body *DetectLiveByFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o DetectLiveByFileRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectLiveByFileRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveByFileRequest", string(data)}, " ")
}
