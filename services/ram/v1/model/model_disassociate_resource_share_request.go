package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateResourceShareRequest Request Object
type DisassociateResourceShareRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *ResourceShareAssociationReqBody `json:"body,omitempty"`
}

func (o DisassociateResourceShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResourceShareRequest struct{}"
	}

	return strings.Join([]string{"DisassociateResourceShareRequest", string(data)}, " ")
}
