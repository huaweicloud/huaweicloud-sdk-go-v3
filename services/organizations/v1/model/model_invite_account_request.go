package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteAccountRequest Request Object
type InviteAccountRequest struct {
	Body *InviteAccountReqBody `json:"body,omitempty"`
}

func (o InviteAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteAccountRequest struct{}"
	}

	return strings.Join([]string{"InviteAccountRequest", string(data)}, " ")
}
