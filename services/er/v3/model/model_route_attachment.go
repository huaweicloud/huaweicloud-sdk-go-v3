package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RouteAttachment struct {

	// 连接关联的资源ID
	ResourceId string `json:"resource_id"`

	// 连接关联的资源类型: - vpc：虚拟私有云 - vpn：vpn网关 - vgw：云专线的虚拟网关 - vpn：vpn网关 - vgw：云专线的虚拟网关 - peering：对等连接，通过云连接CC加载不同区域的企业路由器来创建“对等连接（Peering）”连接
	ResourceType string `json:"resource_type"`

	// 连接ID
	AttachmentId string `json:"attachment_id"`
}

func (o RouteAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteAttachment struct{}"
	}

	return strings.Join([]string{"RouteAttachment", string(data)}, " ")
}
