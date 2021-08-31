package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeGeneralTableResponse struct {
	Result         *GeneralTableResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o RecognizeGeneralTableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTableResponse struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTableResponse", string(data)}, " ")
}
