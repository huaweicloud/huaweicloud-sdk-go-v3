package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用申请资源
type ResourceDto struct {

	// cpu架构类型，不填默认X86
	CpuType *string `json:"cpu_type,omitempty"`

	// cpu申请使用量，取值范围[0.1-128]，单位C，支持一位小数。对于应用，不填默认1C；对于流程和作业，不填默认使用前一级的配置，填值会覆盖更新。覆盖关系：作业->流程->应用
	Cpu *string `json:"cpu,omitempty"`

	// 内存申请使用量，取值范围[0.1-512]，单位G，支持一位小数。对于应用，不填默认1G；对于流程和作业，不填默认使用前一级的配置，填值会覆盖更新。覆盖关系：作业->流程->应用
	Memory *string `json:"memory,omitempty"`

	// gpu架构类型，取值范围 ' '|GPU|D910|D310。对于流程和作业，不填默认使用前一级的配置，填值会覆盖更新。覆盖关系：作业->流程->应用
	GpuType *string `json:"gpu_type,omitempty"`

	// gpu申请使用量，取值范围[0-16]，仅支持整数，D910有特殊约束，申请数量需要是0,1,2,4,8。对于应用，不填默认0；对于流程和作业，不填默认使用前一级的配置，填值会覆盖更新。覆盖关系：作业->流程->应用
	Gpu *string `json:"gpu,omitempty"`
}

func (o ResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDto struct{}"
	}

	return strings.Join([]string{"ResourceDto", string(data)}, " ")
}
