package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RolePrivilegeV5 struct {

	// **参数解释**: 权限信息。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	Privileges []PrivilegeParam `json:"privileges"`
}

func (o RolePrivilegeV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RolePrivilegeV5 struct{}"
	}

	return strings.Join([]string{"RolePrivilegeV5", string(data)}, " ")
}
