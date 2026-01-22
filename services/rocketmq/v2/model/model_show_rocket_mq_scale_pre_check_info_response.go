package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketMqScalePreCheckInfoResponse Response Object
type ShowRocketMqScalePreCheckInfoResponse struct {

	// **参数解释**： 检查项名称。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 检查项状态。 **约束限制**： 不涉及。 **取值范围**： - true：检查项正常。 - false：检查项异常。 **默认取值**： 不涉及。
	Success *bool `json:"success,omitempty"`

	// **参数解释**： 失败原因。   **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Reason *string `json:"reason,omitempty"`

	// **参数解释**： 风险等级。   **约束限制**： 不涉及。 **取值范围**： - low：低风险。 - medium：中风险。 - high：高风险。 **默认取值**： 不涉及。
	Risk           *string `json:"risk,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRocketMqScalePreCheckInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketMqScalePreCheckInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketMqScalePreCheckInfoResponse", string(data)}, " ")
}
