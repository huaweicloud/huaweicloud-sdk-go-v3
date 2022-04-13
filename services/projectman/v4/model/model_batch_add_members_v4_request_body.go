package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAddMembersV4RequestBody struct {
	// 添加的用户信息

	Users []BatchAddMemberRequestV4 `json:"users"`
}

func (o BatchAddMembersV4RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersV4RequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddMembersV4RequestBody", string(data)}, " ")
}
