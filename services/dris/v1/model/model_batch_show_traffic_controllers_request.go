package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowTrafficControllersRequest Request Object
type BatchShowTrafficControllersRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：分页查询时的页码， offset大于等于0，默认取值为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：分页查询时每页显示的记录数，默认值为10，取值范围为0-20的整数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：信号机设备ID，全局唯一。  **取值范围**：长度不超过128，只允许字母、数字、以及_-等字符的组合。
	TrafficControllerId *string `json:"traffic_controller_id,omitempty"`

	// **参数说明**：序列号。  **取值范围**：长度不超过64，只允许字母、数字、以及_等字符的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：设备状态。  **取值范围**： - ONLINE：在线 - OFFLINE：离线 - INITIAL：初始化
	Status *string `json:"status,omitempty"`
}

func (o BatchShowTrafficControllersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowTrafficControllersRequest struct{}"
	}

	return strings.Join([]string{"BatchShowTrafficControllersRequest", string(data)}, " ")
}
