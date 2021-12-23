package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateTagResponse", string(data)}, " ")
}
