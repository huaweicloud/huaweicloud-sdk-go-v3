package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchShowVehiclesRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：分页查询时的页码， offset大于等于0，默认取值为0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：分页查询时每页显示的记录数，默认值为10，取值范围为0-20的整数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：车辆唯一标识符。  **取值范围**：长度不超过128，只允许字母、数字、以及_-等字符的组合。
	VehicleId *string `json:"vehicle_id,omitempty"`

	// **参数说明**：设备状态。  **取值范围**： - ONLINE：在线 - OFFLINE：离线 - INITIAL：初始化
	Status *string `json:"status,omitempty"`
}

func (o BatchShowVehiclesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowVehiclesRequest struct{}"
	}

	return strings.Join([]string{"BatchShowVehiclesRequest", string(data)}, " ")
}
