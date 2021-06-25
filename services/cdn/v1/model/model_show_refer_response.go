package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowReferResponse struct {
	Referer        *RefererRsp `json:"referer,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowReferResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowReferResponse struct{}"
	}

	return strings.Join([]string{"ShowReferResponse", string(data)}, " ")
}
