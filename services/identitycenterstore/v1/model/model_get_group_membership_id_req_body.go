package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetGroupMembershipIdReqBody 获取关联关系唯一标识请求体
type GetGroupMembershipIdReqBody struct {

	// 身份源中IAM身份中心用户组的全局唯一标识符（ID）
	GroupId string `json:"group_id"`

	MemberId *MemberIdDto `json:"member_id"`
}

func (o GetGroupMembershipIdReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetGroupMembershipIdReqBody struct{}"
	}

	return strings.Join([]string{"GetGroupMembershipIdReqBody", string(data)}, " ")
}
