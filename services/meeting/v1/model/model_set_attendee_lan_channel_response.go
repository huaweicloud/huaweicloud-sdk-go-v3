package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAttendeeLanChannelResponse Response Object
type SetAttendeeLanChannelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetAttendeeLanChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAttendeeLanChannelResponse struct{}"
	}

	return strings.Join([]string{"SetAttendeeLanChannelResponse", string(data)}, " ")
}
