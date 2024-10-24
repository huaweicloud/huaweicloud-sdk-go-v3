package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateResourceSharePermissionRequest Request Object
type DisassociateResourceSharePermissionRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *DisassociatePermissionReqBody `json:"body,omitempty"`
}

func (o DisassociateResourceSharePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResourceSharePermissionRequest struct{}"
	}

	return strings.Join([]string{"DisassociateResourceSharePermissionRequest", string(data)}, " ")
}
