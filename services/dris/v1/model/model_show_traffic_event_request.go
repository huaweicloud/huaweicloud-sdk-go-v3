package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrafficEventRequest Request Object
type ShowTrafficEventRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：事件ID，创建事件后获得。方法参见 [新增交通事件](https://support.huaweicloud.com/api-v2x/v2x_04_0048.html)。
	EventId string `json:"event_id"`
}

func (o ShowTrafficEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrafficEventRequest struct{}"
	}

	return strings.Join([]string{"ShowTrafficEventRequest", string(data)}, " ")
}
