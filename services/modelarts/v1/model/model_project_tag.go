package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectTag 项目下训练作业标签，包含key和该key下所有不同的value。
type ProjectTag struct {

	// **参数解释**：标签key。 **取值范围**：1-128字符，支持字母、数字、中文及特殊字符（_.:=+-@），不能以_sys_开头。
	Key *string `json:"key,omitempty"`

	// **参数解释**：该key下出现过的所有不同value列表，已去重。 **取值范围**：0-255字符，支持字母、数字、中文及特殊字符（_.:/=+-@）。
	Values *[]string `json:"values,omitempty"`
}

func (o ProjectTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTag struct{}"
	}

	return strings.Join([]string{"ProjectTag", string(data)}, " ")
}
