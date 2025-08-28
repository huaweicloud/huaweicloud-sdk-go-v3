package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLogtankOption struct {

	// **参数解释**：负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	// **参数解释**：日志组别id，其他（非ELB）服务提供。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LogGroupId string `json:"log_group_id"`

	// **参数解释**：日志订阅主题id，其他（非ELB）服务提供。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LogTopicId string `json:"log_topic_id"`
}

func (o CreateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankOption struct{}"
	}

	return strings.Join([]string{"CreateLogtankOption", string(data)}, " ")
}
