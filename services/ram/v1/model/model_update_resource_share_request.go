package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceShareRequest Request Object
type UpdateResourceShareRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *UpdateResourceShareReqBody `json:"body,omitempty"`
}

func (o UpdateResourceShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceShareRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceShareRequest", string(data)}, " ")
}
