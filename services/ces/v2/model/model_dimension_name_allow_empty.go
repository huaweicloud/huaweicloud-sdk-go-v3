package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimensionNameAllowEmpty **参数解释**： 资源维度。 **约束限制**： 当template_type为0或者不选时，DimensionName必填。当template_type为2时，DimensionName为空。 **取值范围**： 必须以字母开头，多维度用\",\"分割，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32。字符串最大长度为131。 **默认取值**： 不涉及。
type DimensionNameAllowEmpty struct {
}

func (o DimensionNameAllowEmpty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionNameAllowEmpty struct{}"
	}

	return strings.Join([]string{"DimensionNameAllowEmpty", string(data)}, " ")
}
