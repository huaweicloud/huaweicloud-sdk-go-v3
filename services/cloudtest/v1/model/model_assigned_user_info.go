package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssignedUserInfo struct {
	// id信息

	Id *string `json:"id,omitempty"`
	// 名称信息

	Name *string `json:"name,omitempty"`
}

func (o AssignedUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssignedUserInfo struct{}"
	}

	return strings.Join([]string{"AssignedUserInfo", string(data)}, " ")
}
