package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableCorsVpcsRequest Request Object
type ShowAvailableCorsVpcsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowAvailableCorsVpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableCorsVpcsRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableCorsVpcsRequest", string(data)}, " ")
}
