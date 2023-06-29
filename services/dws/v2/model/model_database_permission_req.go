package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabasePermissionReq 数据库权限请求
type DatabasePermissionReq struct {

	// 对象类型 [DATABASE | SCHEMA | TABLE | VIEW | COLUMN | FUNCTION| SEQUENCE | NODEGROUP | ROLE]
	Type string `json:"type"`

	// 是否授权操作
	IsGrant bool `json:"is_grant"`

	// 授权列表 is_grant为true时必填
	GrantList *[]Grant `json:"grant_list,omitempty"`

	// 撤销权限列表 is_grant为false时必填
	RevokeList *[]Revoke `json:"revoke_list,omitempty"`

	// 被授权角色列表
	RoleList []string `json:"role_list"`

	// 权限所属对象列表
	ObjectList *interface{} `json:"object_list"`

	// 撤销权限是否级联撤销 默认 true
	Cascade *bool `json:"cascade,omitempty"`

	// 数据库名称
	Database string `json:"database"`

	// 模式名称
	Schema *string `json:"schema,omitempty"`

	// 表名
	Table *string `json:"table,omitempty"`
}

func (o DatabasePermissionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabasePermissionReq struct{}"
	}

	return strings.Join([]string{"DatabasePermissionReq", string(data)}, " ")
}
