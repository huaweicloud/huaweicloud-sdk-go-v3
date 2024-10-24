package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptResourceShareInvitationRequest Request Object
type AcceptResourceShareInvitationRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源共享邀请的ID。
	ResourceShareInvitationId string `json:"resource_share_invitation_id"`
}

func (o AcceptResourceShareInvitationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptResourceShareInvitationRequest struct{}"
	}

	return strings.Join([]string{"AcceptResourceShareInvitationRequest", string(data)}, " ")
}
