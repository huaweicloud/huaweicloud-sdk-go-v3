package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeGeneralTextResponse struct {
	Result         *GeneralTextResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeGeneralTextResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTextResponse struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTextResponse", string(data)}, " ")
}
