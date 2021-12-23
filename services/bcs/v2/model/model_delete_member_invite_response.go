package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteMemberInviteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMemberInviteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberInviteResponse struct{}"
	}

	return strings.Join([]string{"DeleteMemberInviteResponse", string(data)}, " ")
}
