package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProxyPrivateDnsNameRequest Request Object
type DeleteProxyPrivateDnsNameRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  数据库代理ID，严格匹配UUID规则。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为po01，长度为36个字符。  **默认取值**：  不涉及。
	ProxyId string `json:"proxy_id"`
}

func (o DeleteProxyPrivateDnsNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProxyPrivateDnsNameRequest struct{}"
	}

	return strings.Join([]string{"DeleteProxyPrivateDnsNameRequest", string(data)}, " ")
}
