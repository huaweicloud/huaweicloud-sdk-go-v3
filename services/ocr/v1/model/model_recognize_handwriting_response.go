package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeHandwritingResponse struct {
	Result         *HandwritingResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RecognizeHandwritingResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeHandwritingResponse struct{}"
	}

	return strings.Join([]string{"RecognizeHandwritingResponse", string(data)}, " ")
}
