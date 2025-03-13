package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HighPermission 高权限管理诊断结果。
type HighPermission struct {
	Result *DiagnoseResult `json:"result,omitempty"`

	// 空间管理员用户列表。
	WorkspaceAdmin *string `json:"workspace_admin,omitempty"`

	// 安全管理员用户列表。
	SecurityAdministrator *string `json:"security_administrator,omitempty"`
}

func (o HighPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HighPermission struct{}"
	}

	return strings.Join([]string{"HighPermission", string(data)}, " ")
}
