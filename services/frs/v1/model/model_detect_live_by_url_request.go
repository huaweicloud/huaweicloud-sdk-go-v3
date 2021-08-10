package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectLiveByUrlRequest struct {
	Body *LiveDetectUrlReq `json:"body,omitempty"`
}

func (o DetectLiveByUrlRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectLiveByUrlRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveByUrlRequest", string(data)}, " ")
}
