package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeOperationReq 订阅操作请求，支持项目、用户、用户组的添加或删除订阅,项目、用户、用户组至少存在一个。
type SubscribeOperationReq struct {
	Project *SubscribeOperationReqProject `json:"project,omitempty"`

	// 用户信息列表
	Users *[]CreateSubscribeUserInfo `json:"users,omitempty"`

	// 用户组信息列表
	Usergroups *[]CreateSubscribeUserGroupInfo `json:"usergroups,omitempty"`
}

func (o SubscribeOperationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeOperationReq struct{}"
	}

	return strings.Join([]string{"SubscribeOperationReq", string(data)}, " ")
}
