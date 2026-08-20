package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagFilter 单个标签筛选条件。
type ListTagFilter struct {

	// **参数解释**：标签键。 **约束限制**：   - 必填；   - 长度 1~128；   - 首尾不可为空格；   - 仅支持字母、数字、空格及 `_ . : = + - @`。 **取值范围**：符合标签键命名规范的字符串。 **默认取值**：不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**：标签值列表，与 `key` 组合用于筛选作业。 **约束限制**：   - 非必填；   - 最多 10 个值；   - 单个值长度 0~255；   - 仅支持字母、数字、空格及 `_ . : / = + - @`。 **取值范围**：   - 传具体值：匹配 `key=value` 的作业；   - 不传、传空数组或空字符串：匹配带有该 `key` 的作业（不限 value）。 **默认取值**：不涉及。
	Values *[]string `json:"values,omitempty"`
}

func (o ListTagFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagFilter struct{}"
	}

	return strings.Join([]string{"ListTagFilter", string(data)}, " ")
}
