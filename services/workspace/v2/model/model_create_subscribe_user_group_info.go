package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscribeUserGroupInfo 订阅用户组信息
type CreateSubscribeUserGroupInfo struct {

	// 用户组ID。
	GroupId string `json:"group_id"`
}

func (o CreateSubscribeUserGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscribeUserGroupInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscribeUserGroupInfo", string(data)}, " ")
}
