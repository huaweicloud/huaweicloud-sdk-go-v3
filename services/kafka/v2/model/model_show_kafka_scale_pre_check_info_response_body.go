package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowKafkaScalePreCheckInfoResponseBody struct {

	// **参数解释**： 检查项名称。  **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 检查项状态。 **取值范围**： - true：正常。 - false：异常。
	Success *bool `json:"success,omitempty"`

	// **参数解释**： 失败原因。    **取值范围**： 不涉及。
	Reason *string `json:"reason,omitempty"`

	// **参数解释**： 风险等级。   **取值范围**： - low：低风险。 - medium：中风险。 - high：高风险。
	Risk *string `json:"risk,omitempty"`
}

func (o ShowKafkaScalePreCheckInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaScalePreCheckInfoResponseBody struct{}"
	}

	return strings.Join([]string{"ShowKafkaScalePreCheckInfoResponseBody", string(data)}, " ")
}
