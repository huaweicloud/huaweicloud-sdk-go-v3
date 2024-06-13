package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NpuDetailsDto NPU设备信息
type NpuDetailsDto struct {

	// 昇腾设备ID
	DeviceId *string `json:"device_id,omitempty"`

	// 华为AI加速卡型号，如D310推理卡、D310P推理卡、D910训练卡。
	NpuType *string `json:"npu_type,omitempty"`

	// 昇腾设备产品类型
	ProductName *string `json:"product_name,omitempty"`

	// AI加速卡包含ai核个数
	AiCore *int32 `json:"ai_core,omitempty"`

	// 昇腾设备健康状态
	Health *string `json:"health,omitempty"`

	// 昇腾设备故障信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// NPU使用信息
	UsedInfo *[]NpuUsedInfoDto `json:"used_info,omitempty"`
}

func (o NpuDetailsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NpuDetailsDto struct{}"
	}

	return strings.Join([]string{"NpuDetailsDto", string(data)}, " ")
}
