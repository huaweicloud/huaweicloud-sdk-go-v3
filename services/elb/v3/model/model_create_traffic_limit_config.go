package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrafficLimitConfig **参数解释**：转发策略限速的配置。  **约束限制**：不涉及
type CreateTrafficLimitConfig struct {

	// **参数解释**：转发策略qps限速。  **约束限制**：不涉及  **取值范围**：0-100000，单位：个/秒。0表示不限速。  **默认取值**：不涉及
	Qps *int32 `json:"qps,omitempty"`

	// **参数解释**：对转发策略单源(单个客户端IP)进行限速。  **约束限制**： - quic监听器下转发策略不支持配置单源限速。 - 指定该字段时，赋值可以为0或者为null。 - 如果qps不为0，per_source_ip_qps需要小于qps。  **取值范围**：0-100000，单位：个/秒。0表示不限速。  **默认取值**：不涉及
	PerSourceIpQps *int32 `json:"per_source_ip_qps,omitempty"`

	// **参数解释**：设置当单源qps超限时，允许的局部突增请求数量。超出该限制的请求将返回503。  **约束限制**：不涉及  **取值范围**：0-100000，单位：个/秒。  **默认取值**：不涉及
	Burst *int32 `json:"burst,omitempty"`
}

func (o CreateTrafficLimitConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrafficLimitConfig struct{}"
	}

	return strings.Join([]string{"CreateTrafficLimitConfig", string(data)}, " ")
}
