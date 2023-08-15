package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActivateNodeRequestBody struct {

	// 订单ID，小型轻型设备激活时使用的订单
	OrderId *string `json:"order_id,omitempty"`

	// 订单ID，大型设备使用CPU时激活的订单
	CpuOrderId *string `json:"cpu_order_id,omitempty"`

	// 订单ID，大型设备使用GPU/NPU时激活的订单
	NpuGpuOrderId *string `json:"npu_gpu_order_id,omitempty"`
}

func (o ActivateNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActivateNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ActivateNodeRequestBody", string(data)}, " ")
}
