package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDetailsRequest Request Object
type ShowTaskDetailsRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  任务ID。  **约束限制**：   不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	JobId string `json:"job_id"`

	// **参数解释**：  任务名称。  **约束限制**：   不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	JobName string `json:"job_name"`
}

func (o ShowTaskDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailsRequest", string(data)}, " ")
}
