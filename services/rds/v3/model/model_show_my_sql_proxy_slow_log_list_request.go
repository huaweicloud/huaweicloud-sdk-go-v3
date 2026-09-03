package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMySqlProxySlowLogListRequest Request Object
type ShowMySqlProxySlowLogListRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  数据库代理ID，此参数是数据库代理的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	ProxyId string `json:"proxy_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  查询开始时间，毫秒级时间戳。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**：  查询结束时间，毫秒级时间戳。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime int64 `json:"end_time"`

	// **参数解释**：  每页条数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  10。
	PerPage *int32 `json:"per_page,omitempty"`

	// **参数解释**：  每次查询起始记录ID。首次查询可不传，后续传入返回值中的line_num以获取后续动态加载的数据。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	LineNum *string `json:"line_num,omitempty"`
}

func (o ShowMySqlProxySlowLogListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMySqlProxySlowLogListRequest struct{}"
	}

	return strings.Join([]string{"ShowMySqlProxySlowLogListRequest", string(data)}, " ")
}
