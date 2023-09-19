package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMemberGroupAndStreamRequest Request Object
type ShowMemberGroupAndStreamRequest struct {

	// 成员账号ID
	MemberAccountId string `json:"member_account_id"`
}

func (o ShowMemberGroupAndStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberGroupAndStreamRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberGroupAndStreamRequest", string(data)}, " ")
}
