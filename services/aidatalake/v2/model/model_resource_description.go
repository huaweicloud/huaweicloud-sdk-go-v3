package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceDescription **参数解释**：资源描述信息。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~512个字符。
type ResourceDescription struct {
}

func (o ResourceDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDescription struct{}"
	}

	return strings.Join([]string{"ResourceDescription", string(data)}, " ")
}
