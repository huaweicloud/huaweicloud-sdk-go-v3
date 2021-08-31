package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeBankcardRequest struct {
	Body *BankcardRequestBody `json:"body,omitempty"`
}

func (o RecognizeBankcardRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeBankcardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeBankcardRequest", string(data)}, " ")
}
