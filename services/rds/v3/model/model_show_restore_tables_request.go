package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestoreTablesRequest Request Object
type ShowRestoreTablesRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ShowRestoreTablesRequestBody `json:"body,omitempty"`
}

func (o ShowRestoreTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestoreTablesRequest struct{}"
	}

	return strings.Join([]string{"ShowRestoreTablesRequest", string(data)}, " ")
}
