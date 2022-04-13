package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type MemberRef struct {
	// 后端服务器ID。

	Id string `json:"id"`
}

func (o MemberRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberRef struct{}"
	}

	return strings.Join([]string{"MemberRef", string(data)}, " ")
}
