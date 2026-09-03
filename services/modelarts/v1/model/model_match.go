package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Match 资源名称匹配条件。
type Match struct {

	// **参数解释**：搜索字段名。 **取值范围**：resource_name。
	Key *string `json:"key,omitempty"`

	// **参数解释**：搜索关键词，按资源名称模糊匹配。 **取值范围**：1-255字符。
	Value *string `json:"value,omitempty"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
