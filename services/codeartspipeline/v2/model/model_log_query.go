package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogQuery **参数解释**： 流水线日志查询请求体。 **约束限制**： startOffset 和 endOffset 均设置为 0，则代表查询全量日志。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type LogQuery struct {

	// **参数解释**： 日志起始偏移。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartOffset *int64 `json:"start_offset,omitempty"`

	// **参数解释**： 日志结束偏移。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndOffset *int64 `json:"end_offset,omitempty"`

	// **参数解释**： 最大日志行数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit int64 `json:"limit"`

	// **参数解释**： 排序规则。 **约束限制**： 不涉及。 **取值范围**： - asc：按排序字段升序。 - desc：按排序字段降序 **默认取值**： 不涉及。
	Sort string `json:"sort"`
}

func (o LogQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogQuery struct{}"
	}

	return strings.Join([]string{"LogQuery", string(data)}, " ")
}
