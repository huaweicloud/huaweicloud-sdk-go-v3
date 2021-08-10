package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectLiveByBase64Request struct {
	Body *LiveDetectBase64Req `json:"body,omitempty"`
}

func (o DetectLiveByBase64Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectLiveByBase64Request struct{}"
	}

	return strings.Join([]string{"DetectLiveByBase64Request", string(data)}, " ")
}
