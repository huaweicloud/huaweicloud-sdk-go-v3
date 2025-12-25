package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimensionNameAllowEmpty **参数解释**： 资源维度。     **约束限制**： 事件告警模板DimensionName为空 **取值范围**： 必须以字母开头，多维度用\",\"分隔，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32。目前最大支持4个维度。字符串总最大长度为131。举例：单维度场景：instance_id；多维度场景：instance_id,disk        **默认取值**： 不涉及。
type DimensionNameAllowEmpty struct {
}

func (o DimensionNameAllowEmpty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionNameAllowEmpty struct{}"
	}

	return strings.Join([]string{"DimensionNameAllowEmpty", string(data)}, " ")
}
