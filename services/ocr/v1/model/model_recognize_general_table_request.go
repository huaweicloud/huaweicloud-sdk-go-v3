package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeGeneralTableRequest struct {
	Body *GeneralTableRequestBody `json:"body,omitempty"`
}

func (o RecognizeGeneralTableRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTableRequest struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTableRequest", string(data)}, " ")
}
