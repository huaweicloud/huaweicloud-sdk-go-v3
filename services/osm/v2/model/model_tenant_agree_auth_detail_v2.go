package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TenantAgreeAuthDetailV2 struct {
	// 端口

	Port *int32 `json:"port,omitempty"`
	// 账号

	Account *string `json:"account,omitempty"`
	// 密码

	Password *string `json:"password,omitempty"`
	// 授权详情id

	AuthDetailId int64 `json:"auth_detail_id"`
}

func (o TenantAgreeAuthDetailV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantAgreeAuthDetailV2 struct{}"
	}

	return strings.Join([]string{"TenantAgreeAuthDetailV2", string(data)}, " ")
}
