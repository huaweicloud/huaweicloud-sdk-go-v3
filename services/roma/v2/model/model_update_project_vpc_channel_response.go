package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectVpcChannelResponse Response Object
type UpdateProjectVpcChannelResponse struct {

	// 项目VPC通道列表
	ProjectVpcChannels *[]ProjectVpcChannelInfo `json:"project_vpc_channels,omitempty"`
	HttpStatusCode     int                      `json:"-"`
}

func (o UpdateProjectVpcChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectVpcChannelResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectVpcChannelResponse", string(data)}, " ")
}
