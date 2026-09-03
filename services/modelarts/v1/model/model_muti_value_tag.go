package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MutiValueTag 多值标签筛选条件。
type MutiValueTag struct {

	// **参数解释**：标签key。 **取值范围**：1-128字符，支持字母、数字、中文及特殊字符（_.:=+-@），不能以_sys_开头。
	Key *string `json:"key,omitempty"`

	// **参数解释**：该key下的value列表，多个value之间为OR关系。 **约束限制**：同一key下values不能重复，最多10个。 **取值范围**：0-255字符，支持字母、数字、中文及特殊字符（_.:/=+-@）。
	Values *[]string `json:"values,omitempty"`
}

func (o MutiValueTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MutiValueTag struct{}"
	}

	return strings.Join([]string{"MutiValueTag", string(data)}, " ")
}
