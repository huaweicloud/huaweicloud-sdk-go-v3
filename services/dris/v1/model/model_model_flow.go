package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelFlow **参数说明**：不同车辆类型的流量统计。
type ModelFlow struct {

	// **参数说明**：车辆类型。参考[车辆基本类型](https://support.huaweicloud.com/api-v2x/v2x_04_0162.html)。
	VehicleClass *int32 `json:"vehicle_class,omitempty"`

	// **参数说明**：统计周期内的车辆数。
	Flow *int32 `json:"flow,omitempty"`

	// **参数说明**：车辆平均速度，单位km/h。
	AverageSpeed float32 `json:"average_speed,omitempty"`
}

func (o ModelFlow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelFlow struct{}"
	}

	return strings.Join([]string{"ModelFlow", string(data)}, " ")
}
