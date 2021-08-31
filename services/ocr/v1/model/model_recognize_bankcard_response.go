package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeBankcardResponse struct {
	Result         *BankcardResult `json:"result,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RecognizeBankcardResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeBankcardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeBankcardResponse", string(data)}, " ")
}
