package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectLiveFaceByBase64Request struct {
	Body *LiveDetectFaceBase64Req `json:"body,omitempty"`
}

func (o DetectLiveFaceByBase64Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByBase64Request struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByBase64Request", string(data)}, " ")
}
