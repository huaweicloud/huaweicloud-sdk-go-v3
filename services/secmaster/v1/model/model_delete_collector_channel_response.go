package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCollectorChannelResponse Response Object
type DeleteCollectorChannelResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteCollectorChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCollectorChannelResponse struct{}"
	}

	return strings.Join([]string{"DeleteCollectorChannelResponse", string(data)}, " ")
}
