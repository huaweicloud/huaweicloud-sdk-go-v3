package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PermissionSetMemberBatchCreateDto struct {

	// 权限集成员创建参数列表
	Members *[]PermissionSetMemberCreateDto2 `json:"members,omitempty"`

	// 是否自动触发同步, 默认false
	AutoSync *bool `json:"auto_sync,omitempty"`
}

func (o PermissionSetMemberBatchCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetMemberBatchCreateDto struct{}"
	}

	return strings.Join([]string{"PermissionSetMemberBatchCreateDto", string(data)}, " ")
}
