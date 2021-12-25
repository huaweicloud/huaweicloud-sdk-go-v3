package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateProjectVpcChannelSyncsResponse struct {
	// 项目VPC通道列表

	ProjectVpcChannels *[]ProjectVpcChannelInfo `json:"project_vpc_channels,omitempty"`
	HttpStatusCode     int                      `json:"-"`
}

func (o CreateProjectVpcChannelSyncsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectVpcChannelSyncsResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectVpcChannelSyncsResponse", string(data)}, " ")
}
