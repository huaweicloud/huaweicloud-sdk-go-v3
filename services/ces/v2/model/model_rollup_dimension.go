package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollupDimension **参数解释** 聚合维度 **约束限制** 不涉及 **取值范围** 长度为[1,32]个字符 **默认取值** 不涉及
type RollupDimension struct {
}

func (o RollupDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollupDimension struct{}"
	}

	return strings.Join([]string{"RollupDimension", string(data)}, " ")
}
