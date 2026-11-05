package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventUpdate struct {

	// **参数解释**：计划执行开始时间，格式为yyyy-MM-ddTHH:mm:ssZ。 **约束限制**：不涉及。 **取值范围**：大于当前时间。 **默认取值**：不填表示立即执行。
	NotBefore *string `json:"notBefore,omitempty"`

	RedeployConfig *RedeployConfig `json:"redeployConfig,omitempty"`

	// **参数解释**：授权节点ID列表。授权重部处理时可选填,支持填入单个故障节点ID。 **约束限制**：不涉及。 **取值范围**：计划事件故障节点列表。 **默认取值**：不填表示授权所有节点。
	InstanceId *[]string `json:"instanceId,omitempty"`
}

func (o EventUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventUpdate struct{}"
	}

	return strings.Join([]string{"EventUpdate", string(data)}, " ")
}
