package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcAttachmentDetails VPC类型连接
type VpcAttachmentDetails struct {

	// VPC连接ID
	Id string `json:"id"`

	// VPC连接名称
	Name string `json:"name"`

	// VPC id
	VpcId string `json:"vpc_id"`

	// VPC子网id
	VirsubnetId string `json:"virsubnet_id"`

	// 默认为false,当设置true时，会自动为VPC配置10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16三条路由，下一跳指向企业路由器。
	AutoCreateVpcRoutes *bool `json:"auto_create_vpc_routes,omitempty"`

	// VPC连接状态:pending|available|modifying|deleting|deleted|failed|initiating_request|rejected|pending_acceptance
	State string `json:"state"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`

	// VPC连接描述信息
	Description *string `json:"description,omitempty"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// vpc所属项目ID
	VpcProjectId *string `json:"vpc_project_id,omitempty"`
}

func (o VpcAttachmentDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcAttachmentDetails struct{}"
	}

	return strings.Join([]string{"VpcAttachmentDetails", string(data)}, " ")
}
