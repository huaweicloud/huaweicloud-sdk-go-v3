package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateTaurusSlowProcessThresholdRequest Request Object
type BatchUpdateTaurusSlowProcessThresholdRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *BatchModifySlowProcessThresholdRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateTaurusSlowProcessThresholdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateTaurusSlowProcessThresholdRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateTaurusSlowProcessThresholdRequest", string(data)}, " ")
}
