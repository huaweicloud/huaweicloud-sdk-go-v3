package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeAutoClassificationRequest struct {
	Body *AutoClassificationRequestBody `json:"body,omitempty"`
}

func (o RecognizeAutoClassificationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeAutoClassificationRequest struct{}"
	}

	return strings.Join([]string{"RecognizeAutoClassificationRequest", string(data)}, " ")
}
