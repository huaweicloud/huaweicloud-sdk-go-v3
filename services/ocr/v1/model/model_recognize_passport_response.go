package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizePassportResponse struct {
	Result         *PassportResult `json:"result,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RecognizePassportResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizePassportResponse struct{}"
	}

	return strings.Join([]string{"RecognizePassportResponse", string(data)}, " ")
}
