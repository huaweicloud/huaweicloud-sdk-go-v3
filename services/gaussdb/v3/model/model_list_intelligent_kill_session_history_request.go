package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIntelligentKillSessionHistoryRequest Request Object
type ListIntelligentKillSessionHistoryRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  节点ID，此参数是节点的唯一标识。  获取方法请参见[查询实例详情](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlInstanceInfoUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**：  所需查询的起始时间戳。  **约束限制**：  \"start_time\"参数值需要小于\"end_time\"参数值。  **取值范围**：  格式为UNIX时间戳，单位是毫秒，时区为UTC标准时区。传参时需要将对应时区的时间转为标准时区对应的时间戳，比如，北京时区的时间点需要-8h后再转为时间戳。  **默认取值**：  不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**：  所需查询的结束时间戳。  **约束限制**：  \"end_time\"参数值需要大于\"start_time\"参数值。  **取值范围**：  格式为UNIX时间戳，单位是毫秒，时区为UTC标准时区。传参时需要将对应时区的时间转为标准时区对应的时间戳，比如，北京时区的时间点需要-8h后再转为时间戳。  **默认取值**：  不涉及。
	EndTime int64 `json:"end_time"`

	// **参数解释**：  查询记录数。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  1-100。  **默认取值**：  100。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  >=0。  **默认取值**：  0。
	Offset *string `json:"offset,omitempty"`
}

func (o ListIntelligentKillSessionHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIntelligentKillSessionHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListIntelligentKillSessionHistoryRequest", string(data)}, " ")
}
