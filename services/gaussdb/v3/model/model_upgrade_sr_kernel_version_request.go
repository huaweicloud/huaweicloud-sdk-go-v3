package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeSrKernelVersionRequest Request Object
type UpgradeSrKernelVersionRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us。
	XLanguage string `json:"X-Language"`

	// **参数解释**：  StarRocks实例ID，严格匹配UUID规则。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，且长度为36个字符。  **默认值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  application/json。
	ContentType string `json:"Content-Type"`

	Body *UpgradeSrKernelVersionRequestV3 `json:"body,omitempty"`
}

func (o UpgradeSrKernelVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeSrKernelVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeSrKernelVersionRequest", string(data)}, " ")
}
