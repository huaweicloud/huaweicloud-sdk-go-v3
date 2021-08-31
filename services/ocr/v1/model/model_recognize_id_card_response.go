package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeIdCardResponse struct {
	Result         *IdCardResult `json:"result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o RecognizeIdCardResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeIdCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeIdCardResponse", string(data)}, " ")
}
