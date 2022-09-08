package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RouteAttachment struct {

	// 连接关联的资源ID
	ResourceId string `json:"resource_id"`

	// 连接关联的资源类型：vpc|vgw|vpn|peering|gdgw
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
