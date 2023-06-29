package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMemberInviteResponse Response Object
type DeleteMemberInviteResponse struct {

	// 请求成功的结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMemberInviteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberInviteResponse struct{}"
	}

	return strings.Join([]string{"DeleteMemberInviteResponse", string(data)}, " ")
}
