package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowRefererResponse struct {
	Referer        *RefererRsp `json:"referer,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowRefererResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRefererResponse struct{}"
	}

	return strings.Join([]string{"ShowRefererResponse", string(data)}, " ")
}
