package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventDimensionName **参数解释** 事件维度名，多个维度按字母序逗号分开 **约束限制** 不涉及 **取值范围** 字符串长度[0,131] **默认取值** 不涉及
type EventDimensionName struct {
}

func (o EventDimensionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDimensionName struct{}"
	}

	return strings.Join([]string{"EventDimensionName", string(data)}, " ")
}
