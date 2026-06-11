package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Scope **参数解释**: 是否选择所有主机 **约束限制**: 不涉及 **取值范围**: - true：是 - false：否 **默认取值**: false
type Scope struct {
}

func (o Scope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scope struct{}"
	}

	return strings.Join([]string{"Scope", string(data)}, " ")
}
