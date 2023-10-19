package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EntityPrivilegeInfo 实体特权信息
type EntityPrivilegeInfo struct {

	// 特权列表
	Privileges *[]string `json:"privileges,omitempty"`

	// 继承特权列表
	InheritPrivileges *[]string `json:"inherit_privileges,omitempty"`
}

func (o EntityPrivilegeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityPrivilegeInfo struct{}"
	}

	return strings.Join([]string{"EntityPrivilegeInfo", string(data)}, " ")
}
