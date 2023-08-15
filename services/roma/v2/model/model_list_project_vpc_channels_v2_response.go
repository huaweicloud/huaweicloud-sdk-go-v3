package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectVpcChannelsV2Response Response Object
type ListProjectVpcChannelsV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的实例VPC通道列表
	ProjectVpcChannels *[]ProjectVpcChannelInfo `json:"project_vpc_channels,omitempty"`
	HttpStatusCode     int                      `json:"-"`
}

func (o ListProjectVpcChannelsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectVpcChannelsV2Response struct{}"
	}

	return strings.Join([]string{"ListProjectVpcChannelsV2Response", string(data)}, " ")
}
