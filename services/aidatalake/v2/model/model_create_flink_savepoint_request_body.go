package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSavepointRequestBody **参数解释**：创建Flink Savepoint的请求参数。 **约束限制**：不涉及。
type CreateFlinkSavepointRequestBody struct {

	// **参数解释**：作业快照的描述信息。 **约束限制**：不涉及。 **取值范围**：0~512个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	Format *FlinkSavepointFormat `json:"format,omitempty"`

	// **参数解释**：是否启用Drain模式。启用后，触发Savepoint前将先发送Watermark到所有算子以触发窗口计算完成，然后再保存状态快照。 **约束限制**：不涉及。 **取值范围**：   - true：启用Drain模式。   - false：不启用Drain模式。 **默认取值**：false。
	DrainModeEnabled *bool `json:"drain_mode_enabled,omitempty"`
}

func (o CreateFlinkSavepointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSavepointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkSavepointRequestBody", string(data)}, " ")
}
