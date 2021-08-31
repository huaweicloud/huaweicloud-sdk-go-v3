package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeWebImageResponse struct {
	Result         *WebImageResult `json:"result,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RecognizeWebImageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeWebImageResponse struct{}"
	}

	return strings.Join([]string{"RecognizeWebImageResponse", string(data)}, " ")
}
