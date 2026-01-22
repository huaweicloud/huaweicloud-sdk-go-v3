package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisStackForRocketMqResponse Response Object
type ShowDiagnosisStackForRocketMqResponse struct {

	// **参数解释**： 线程名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ThreadName *string `json:"thread_name,omitempty"`

	// **参数解释**： 客户端的栈信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Stack          *string `json:"stack,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDiagnosisStackForRocketMqResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisStackForRocketMqResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisStackForRocketMqResponse", string(data)}, " ")
}
