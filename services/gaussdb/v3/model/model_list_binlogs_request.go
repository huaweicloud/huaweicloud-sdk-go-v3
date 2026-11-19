package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBinlogsRequest Request Object
type ListBinlogsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  ≥0。  **默认取值**：  0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。  **约束限制**：  不涉及。  **取值范围**：  - 10：10条。 - 20：20条。 - 50：50条。  **默认取值**：  10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  开始时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **约束限制**：  不得晚于end_time。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**：  结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **约束限制**：  必须晚于或等于begin_time。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListBinlogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBinlogsRequest struct{}"
	}

	return strings.Join([]string{"ListBinlogsRequest", string(data)}, " ")
}
