package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLogtankOption struct {

	// **参数解释**：日志组别id，其他（非ELB）服务提供。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释**：日志订阅主题id，其他（非ELB）服务提供。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LogTopicId *string `json:"log_topic_id,omitempty"`
}

func (o UpdateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankOption struct{}"
	}

	return strings.Join([]string{"UpdateLogtankOption", string(data)}, " ")
}
