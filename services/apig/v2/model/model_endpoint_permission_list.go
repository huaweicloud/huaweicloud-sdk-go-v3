package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointPermissionList struct {

	// 白名单记录列表。每个白名单记录的格式为iam:domain::授权账号ID。  其中，授权账号ID是长度为32的字符串，只包含英文字母（a-f）或数字；也可为*，表示允许全部用户连接。
	Permissions []string `json:"permissions"`
}

func (o EndpointPermissionList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointPermissionList struct{}"
	}

	return strings.Join([]string{"EndpointPermissionList", string(data)}, " ")
}
