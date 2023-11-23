package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StackResource 资源栈
type StackResource struct {

	// 物理资源id
	PhysicalResourceId *string `json:"physical_resource_id,omitempty"`

	// 物理资源名称
	PhysicalResourceName *string `json:"physical_resource_name,omitempty"`

	// 逻辑资源名称
	LogicalResourceName *string `json:"logical_resource_name,omitempty"`

	// 逻辑资源类型
	LogicalResourceType *string `json:"logical_resource_type,omitempty"`

	// 资源状态
	ResourceStatus *string `json:"resource_status,omitempty"`

	// 状态信息
	StatusMessage *string `json:"status_message,omitempty"`

	// 超链接地址
	Href *string `json:"href,omitempty"`

	// 云服务名称
	DisplayName *string `json:"display_name,omitempty"`
}

func (o StackResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackResource struct{}"
	}

	return strings.Join([]string{"StackResource", string(data)}, " ")
}
