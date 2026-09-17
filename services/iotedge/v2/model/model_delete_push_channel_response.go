package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePushChannelResponse Response Object
type DeletePushChannelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePushChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePushChannelResponse struct{}"
	}

	return strings.Join([]string{"DeletePushChannelResponse", string(data)}, " ")
}
