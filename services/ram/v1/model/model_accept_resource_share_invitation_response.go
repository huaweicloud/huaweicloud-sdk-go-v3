package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptResourceShareInvitationResponse Response Object
type AcceptResourceShareInvitationResponse struct {
	ResourceShareInvitation *ResourceShareInvitation `json:"resource_share_invitation,omitempty"`
	HttpStatusCode          int                      `json:"-"`
}

func (o AcceptResourceShareInvitationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptResourceShareInvitationResponse struct{}"
	}

	return strings.Join([]string{"AcceptResourceShareInvitationResponse", string(data)}, " ")
}
