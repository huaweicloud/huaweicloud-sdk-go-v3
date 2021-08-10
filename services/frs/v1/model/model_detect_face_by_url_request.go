package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectFaceByUrlRequest struct {
	Body *FaceDetectUrlReq `json:"body,omitempty"`
}

func (o DetectFaceByUrlRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectFaceByUrlRequest struct{}"
	}

	return strings.Join([]string{"DetectFaceByUrlRequest", string(data)}, " ")
}
