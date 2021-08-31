package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeVinResponse struct {
	Result         *VinResult `json:"result,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o RecognizeVinResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeVinResponse struct{}"
	}

	return strings.Join([]string{"RecognizeVinResponse", string(data)}, " ")
}
