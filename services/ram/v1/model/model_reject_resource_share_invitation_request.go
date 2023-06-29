package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RejectResourceShareInvitationRequest Request Object
type RejectResourceShareInvitationRequest struct {

	// 资源共享邀请的ID。
	ResourceShareInvitationId string `json:"resource_share_invitation_id"`
}

func (o RejectResourceShareInvitationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectResourceShareInvitationRequest struct{}"
	}

	return strings.Join([]string{"RejectResourceShareInvitationRequest", string(data)}, " ")
}
