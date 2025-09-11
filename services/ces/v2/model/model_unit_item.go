package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnitItem **参数解释** 单位 **约束限制** 不涉及 **取值范围** 长度为[0,32]个字符 **默认取值** 不涉及
type UnitItem struct {
}

func (o UnitItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnitItem struct{}"
	}

	return strings.Join([]string{"UnitItem", string(data)}, " ")
}
