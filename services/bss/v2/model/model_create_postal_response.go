package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePostalResponse struct {
	// 邮寄地址ID。

	AddressId      *string `json:"address_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostalResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePostalResponse struct{}"
	}

	return strings.Join([]string{"CreatePostalResponse", string(data)}, " ")
}
