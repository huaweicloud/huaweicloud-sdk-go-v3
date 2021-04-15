package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunCelebrityRecognitionRequest struct {
	Body *CelebrityRecognitionReq `json:"body,omitempty"`
}

func (o RunCelebrityRecognitionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunCelebrityRecognitionRequest struct{}"
	}

	return strings.Join([]string{"RunCelebrityRecognitionRequest", string(data)}, " ")
}
