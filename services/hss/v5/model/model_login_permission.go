package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginPermission **参数解释**: 是否允许登录 **约束限制**: 不涉及 **取值范围**: - true：是 - false：否  **默认取值**: 不涉及
type LoginPermission struct {
}

func (o LoginPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginPermission struct{}"
	}

	return strings.Join([]string{"LoginPermission", string(data)}, " ")
}
