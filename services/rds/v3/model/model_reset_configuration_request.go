package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetConfigurationRequest Request Object
type ResetConfigurationRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  参数模板id。  **约束限制**：  只允许填写自定义参数模板。
	ConfigId string `json:"config_id"`
}

func (o ResetConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ResetConfigurationRequest", string(data)}, " ")
}
