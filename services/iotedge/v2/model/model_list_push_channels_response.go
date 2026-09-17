package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPushChannelsResponse Response Object
type ListPushChannelsResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 通道信息
	Channels       *[]ChannelDto `json:"channels,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListPushChannelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPushChannelsResponse struct{}"
	}

	return strings.Join([]string{"ListPushChannelsResponse", string(data)}, " ")
}
