package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AcceptResourceShareInvitationRequest struct {

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
