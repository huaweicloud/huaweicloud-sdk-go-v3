package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NpuUsedInfoDto NPU设备使用详情
type NpuUsedInfoDto struct {

	// 模块名称
	ModuleId *string `json:"module_id,omitempty"`

	// 模块使用AI核的个数
	UsedAiCoreNum *int32 `json:"used_ai_core_num,omitempty"`

	// 模块使用NPU芯片中的cpu核数
	UsedCpuCoreNum *int32 `json:"used_cpu_core_num,omitempty"`
}

func (o NpuUsedInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NpuUsedInfoDto struct{}"
	}

	return strings.Join([]string{"NpuUsedInfoDto", string(data)}, " ")
}
