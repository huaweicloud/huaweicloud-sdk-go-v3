package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentResponse 连接基本信息详情
type AttachmentResponse struct {

	// 连接名字
	Name string `json:"name"`

	// 连接ID
	Id string `json:"id"`

	// 描述信息
	Description string `json:"description"`

	// 连接状态:pending|available|modifying|deleting|deleted|failed|pending_acceptance|rejected|initiating_request|freezed
	State string `json:"state"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 企业路由器关联tag
	Tags *[]Tag `json:"tags,omitempty"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 内部连接关联的资源ID
	ResourceId string `json:"resource_id"`

	// 内部连接关联的资源类型: - vgw：云专线的虚拟网关 - vpn：vpn网关 - peering：对等连接，通过云连接CC加载不同区域的企业路由器来创建“对等连接（Peering）”连接 - vpc：虚拟私有云 -
	ResourceType string `json:"resource_type"`

	// 资源所属项目ID
	ResourceProjectId *string `json:"resource_project_id,omitempty"`
}

func (o AttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentResponse struct{}"
	}

	return strings.Join([]string{"AttachmentResponse", string(data)}, " ")
}
