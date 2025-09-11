package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Type **参数解释**： 告警策略类型，已废弃，不推荐使用。 **约束限制**： 不涉及。 **取值范围**： 只能为auto。          **默认取值**： 不涉及。
type Type struct {
}

func (o Type) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Type struct{}"
	}

	return strings.Join([]string{"Type", string(data)}, " ")
}
