package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsMemberInGroupsReqBody 判断用户是否在用户组请求体
type IsMemberInGroupsReqBody struct {

	// 用户组标识符列表
	GroupIds []string `json:"group_ids"`

	MemberId *MemberIdDto `json:"member_id"`
}

func (o IsMemberInGroupsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsMemberInGroupsReqBody struct{}"
	}

	return strings.Join([]string{"IsMemberInGroupsReqBody", string(data)}, " ")
}
