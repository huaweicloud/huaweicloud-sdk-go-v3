package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchShowVehiclesResponse struct {

	// **参数说明**：返回车辆的总体数量。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：数据列表。
	Vehicles       *[]VehicleDto `json:"vehicles,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o BatchShowVehiclesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowVehiclesResponse struct{}"
	}

	return strings.Join([]string{"BatchShowVehiclesResponse", string(data)}, " ")
}
