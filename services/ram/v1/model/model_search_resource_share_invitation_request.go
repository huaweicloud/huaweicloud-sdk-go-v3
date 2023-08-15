package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceShareInvitationRequest Request Object
type SearchResourceShareInvitationRequest struct {
	Body *SearchResourceShareInvitationReqBody `json:"body,omitempty"`
}

func (o SearchResourceShareInvitationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareInvitationRequest struct{}"
	}

	return strings.Join([]string{"SearchResourceShareInvitationRequest", string(data)}, " ")
}
