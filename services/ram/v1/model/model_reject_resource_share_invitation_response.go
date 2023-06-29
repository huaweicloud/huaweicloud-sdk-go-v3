package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RejectResourceShareInvitationResponse Response Object
type RejectResourceShareInvitationResponse struct {
	ResourceShareInvitation *ResourceShareInvitation `json:"resource_share_invitation,omitempty"`
	HttpStatusCode          int                      `json:"-"`
}

func (o RejectResourceShareInvitationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectResourceShareInvitationResponse struct{}"
	}

	return strings.Join([]string{"RejectResourceShareInvitationResponse", string(data)}, " ")
}
