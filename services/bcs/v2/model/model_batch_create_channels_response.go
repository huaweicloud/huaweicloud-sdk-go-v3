package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateChannelsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateChannelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateChannelsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateChannelsResponse", string(data)}, " ")
}
