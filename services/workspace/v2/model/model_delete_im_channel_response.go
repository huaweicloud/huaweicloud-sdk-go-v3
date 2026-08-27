package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImChannelResponse Response Object
type DeleteImChannelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteImChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImChannelResponse struct{}"
	}

	return strings.Join([]string{"DeleteImChannelResponse", string(data)}, " ")
}
