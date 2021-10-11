package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAccessCodeResponse struct {
	// 设备接入码

	AccessCode     *string `json:"access_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAccessCodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAccessCodeResponse struct{}"
	}

	return strings.Join([]string{"CreateAccessCodeResponse", string(data)}, " ")
}
