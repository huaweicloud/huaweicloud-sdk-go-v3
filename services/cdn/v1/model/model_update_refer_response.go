package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateReferResponse struct {
	Referer        *RefererRsp `json:"referer,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateReferResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateReferResponse struct{}"
	}

	return strings.Join([]string{"UpdateReferResponse", string(data)}, " ")
}
