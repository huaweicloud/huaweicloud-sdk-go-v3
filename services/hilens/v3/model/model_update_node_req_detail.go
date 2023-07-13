package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNodeReqDetail struct {

	// 设备描述，最大长度255，不允许^, ~, ＃, $, %, &, *, <, >, (, ), [, ], {, }, ', \", \\
	Description *string `json:"description,omitempty"`

	// 设备日志配置
	LogConfigs *[]LogConfig `json:"log_configs,omitempty"`

	// 设备标签，标签个数最多为20个
	Tags *[]TagObject `json:"tags,omitempty"`

	// 事件有效时间(单位：分钟)
	EventValidityPeriod *int32 `json:"event_validity_period,omitempty"`

	// 是否开启gpu
	EnableGpu *bool `json:"enable_gpu,omitempty"`

	// 是否开启npu
	EnableNpu *bool `json:"enable_npu,omitempty"`

	// npu类型，如果选择开启npu, 可设置类型Ascend 310/ Ascend 710, 如果选择开启gpu，请设置值为null。
	NpuType *string `json:"npu_type,omitempty"`
}

func (o UpdateNodeReqDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeReqDetail struct{}"
	}

	return strings.Join([]string{"UpdateNodeReqDetail", string(data)}, " ")
}
