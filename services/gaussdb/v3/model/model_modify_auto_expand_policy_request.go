package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyAutoExpandPolicyRequest Request Object
type ModifyAutoExpandPolicyRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *ModifyAutoExpandPolicyReq `json:"body,omitempty"`
}

func (o ModifyAutoExpandPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAutoExpandPolicyRequest struct{}"
	}

	return strings.Join([]string{"ModifyAutoExpandPolicyRequest", string(data)}, " ")
}
