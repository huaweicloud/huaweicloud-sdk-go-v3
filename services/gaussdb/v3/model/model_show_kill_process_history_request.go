package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessHistoryRequest Request Object
type ShowKillProcessHistoryRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  Kill会话的开始时间。  **约束限制**：  - 格式为UNIX时间戳，单位是毫秒。 - start_time要小于end_time。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：    Kill会话的结束时间。    **约束限制**：    - 格式为UNIX时间戳，单位是毫秒。 - end_time要大于start_time。   **取值范围**：    不涉及。  **默认取值**：    不涉及。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**：    当前页数。    **约束限制**：    必须为整数，不能为负数。    **取值范围**：    ≥0。  **默认取值**：    1。
	PageNum *int32 `json:"page_num,omitempty"`

	// **参数解释**：    每页大小。    **约束限制**：    必须为整数，不能为负数。    **取值范围**：    1-100。  **默认取值**：    10。
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ShowKillProcessHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowKillProcessHistoryRequest", string(data)}, " ")
}
