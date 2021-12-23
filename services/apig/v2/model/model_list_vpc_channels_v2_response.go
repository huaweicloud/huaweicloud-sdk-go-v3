package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListVpcChannelsV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次查询到的VPC通道列表

	VpcChannels    *[]VpcChannelInfo `json:"vpc_channels,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListVpcChannelsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcChannelsV2Response struct{}"
	}

	return strings.Join([]string{"ListVpcChannelsV2Response", string(data)}, " ")
}
