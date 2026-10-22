package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComputeDescription **参数解释**：资源描述信息。 **约束限制**：不涉及。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~512个字符。 **默认取值**：不涉及。
type ComputeDescription struct {
}

func (o ComputeDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeDescription struct{}"
	}

	return strings.Join([]string{"ComputeDescription", string(data)}, " ")
}
