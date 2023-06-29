package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectVpcChannelResponse Response Object
type CreateProjectVpcChannelResponse struct {

	// 项目VPC通道列表
	ProjectVpcChannels *[]ProjectVpcChannelInfo `json:"project_vpc_channels,omitempty"`
	HttpStatusCode     int                      `json:"-"`
}

func (o CreateProjectVpcChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectVpcChannelResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectVpcChannelResponse", string(data)}, " ")
}
