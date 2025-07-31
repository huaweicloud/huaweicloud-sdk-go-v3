package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RootPermission **参数解释**： 管理员权限 **取值范围**： - true：是 - false：否
type RootPermission struct {
}

func (o RootPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RootPermission struct{}"
	}

	return strings.Join([]string{"RootPermission", string(data)}, " ")
}
