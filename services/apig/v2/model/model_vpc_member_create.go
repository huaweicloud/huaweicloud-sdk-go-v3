package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcMemberCreate struct {
	// 后端实例列表

	Members []MemberInfo `json:"members"`
}

func (o VpcMemberCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcMemberCreate struct{}"
	}

	return strings.Join([]string{"VpcMemberCreate", string(data)}, " ")
}
