package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBlockchainChannelsResponse struct {
	// 通道信息列表

	Channels       *[]Channel `json:"channels,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListBlockchainChannelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlockchainChannelsResponse struct{}"
	}

	return strings.Join([]string{"ListBlockchainChannelsResponse", string(data)}, " ")
}
