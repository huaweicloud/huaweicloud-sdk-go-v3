package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyFlavorsByAzCodeRequest Request Object
type ShowProxyFlavorsByAzCodeRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  可用区。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	AzCodes string `json:"az_codes"`

	// **参数解释**：  代理引擎名称。  **约束限制**：  不涉及。  **取值范围**：  taurusproxy。  **默认取值**：  不涉及。
	ProxyEngineName string `json:"proxy_engine_name"`
}

func (o ShowProxyFlavorsByAzCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyFlavorsByAzCodeRequest struct{}"
	}

	return strings.Join([]string{"ShowProxyFlavorsByAzCodeRequest", string(data)}, " ")
}
