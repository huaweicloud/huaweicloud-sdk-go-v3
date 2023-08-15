package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowTrafficControllersResponse Response Object
type BatchShowTrafficControllersResponse struct {

	// **参数说明**：返回信号机设备的总体数量。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：数据列表。
	TrafficControllerDevices *[]TrafficControllerDto `json:"traffic_controller_devices,omitempty"`
	HttpStatusCode           int                     `json:"-"`
}

func (o BatchShowTrafficControllersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowTrafficControllersResponse struct{}"
	}

	return strings.Join([]string{"BatchShowTrafficControllersResponse", string(data)}, " ")
}
