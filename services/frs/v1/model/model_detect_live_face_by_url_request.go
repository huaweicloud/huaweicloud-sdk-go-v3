package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectLiveFaceByUrlRequest struct {
	Body *LiveDetectFaceUrlReq `json:"body,omitempty"`
}

func (o DetectLiveFaceByUrlRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByUrlRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByUrlRequest", string(data)}, " ")
}
