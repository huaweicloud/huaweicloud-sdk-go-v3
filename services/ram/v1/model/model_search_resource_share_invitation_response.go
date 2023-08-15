package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceShareInvitationResponse Response Object
type SearchResourceShareInvitationResponse struct {
	ResourceShareInvitations *[]ResourceShareInvitation `json:"resource_share_invitations,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SearchResourceShareInvitationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareInvitationResponse struct{}"
	}

	return strings.Join([]string{"SearchResourceShareInvitationResponse", string(data)}, " ")
}
