package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteChannelResponse Response Object
type DeleteChannelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteChannelResponse struct{}"
	}

	return strings.Join([]string{"DeleteChannelResponse", string(data)}, " ")
}
