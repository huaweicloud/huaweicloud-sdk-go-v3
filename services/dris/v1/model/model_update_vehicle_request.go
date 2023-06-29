package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVehicleRequest Request Object
type UpdateVehicleRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：车辆唯一标识符。  **取值范围**：长度不超过128，只允许字母、数字、以及_-等字符的组合。
	VehicleId string `json:"vehicle_id"`

	Body *ModifyVehicleRequestDto `json:"body,omitempty"`
}

func (o UpdateVehicleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVehicleRequest struct{}"
	}

	return strings.Join([]string{"UpdateVehicleRequest", string(data)}, " ")
}
