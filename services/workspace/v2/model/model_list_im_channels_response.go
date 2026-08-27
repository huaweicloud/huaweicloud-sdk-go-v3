package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImChannelsResponse Response Object
type ListImChannelsResponse struct {

	// IM 通道配置列表
	ImChannels     *[]ImChannelItem `json:"im_channels,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListImChannelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImChannelsResponse struct{}"
	}

	return strings.Join([]string{"ListImChannelsResponse", string(data)}, " ")
}
