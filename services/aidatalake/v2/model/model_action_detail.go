package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionDetail **参数解释**：权限策略中的action详情，包含action操作和对应的描述。
type ActionDetail struct {

	// **参数解释**：权限操作内容。 **取值范围**：只能由英文字母、数字、冒号及特殊字符组成，长度0~128个字符。
	Action string `json:"action"`

	// **参数解释**：权限操作的描述，说明该操作的具体含义。 **取值范围**：支持中文字符、英文字母、数字及特殊字符，长度0~1024个字符。
	ActionDescription string `json:"action_description"`
}

func (o ActionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionDetail struct{}"
	}

	return strings.Join([]string{"ActionDetail", string(data)}, " ")
}
