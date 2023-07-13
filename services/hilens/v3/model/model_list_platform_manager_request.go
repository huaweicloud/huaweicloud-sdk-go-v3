package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlatformManagerRequest Request Object
type ListPlatformManagerRequest struct {

	// 订单Id，可以根据订单Id查询
	Id *string `json:"id,omitempty"`

	// 设备类别：lite_device轻量型设备，small_device小型设备，large_device大型设备，large_device_cpu CPU大型设备， large_device_gpu_npu CPU,NPU大型设备
	DeviceType *string `json:"device_type,omitempty"`

	// 运行服务费类别，专业版为running_service，标准版为development_service
	Type *string `json:"type,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPlatformManagerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlatformManagerRequest struct{}"
	}

	return strings.Join([]string{"ListPlatformManagerRequest", string(data)}, " ")
}
