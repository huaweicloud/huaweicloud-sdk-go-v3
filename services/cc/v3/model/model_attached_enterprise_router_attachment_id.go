package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachedEnterpriseRouterAttachmentId 被挂载的企业路由器的连接ID。
type AttachedEnterpriseRouterAttachmentId struct {

	// 被挂载的企业路由器的连接ID。
	AttachedErAttachmentId *string `json:"attached_er_attachment_id,omitempty"`
}

func (o AttachedEnterpriseRouterAttachmentId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedEnterpriseRouterAttachmentId struct{}"
	}

	return strings.Join([]string{"AttachedEnterpriseRouterAttachmentId", string(data)}, " ")
}
