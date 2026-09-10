package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelServiceConfigCustomSpec 自定义规格配置
type ModelServiceConfigCustomSpec struct {

	// CPU核数
	Cpu *string `json:"cpu,omitempty"`

	// GPU个数
	Gpu *string `json:"gpu,omitempty"`

	// 昇腾芯片个数
	Ascend *string `json:"ascend,omitempty"`

	// 内存大小
	Memory *string `json:"memory,omitempty"`
}

func (o ModelServiceConfigCustomSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelServiceConfigCustomSpec struct{}"
	}

	return strings.Join([]string{"ModelServiceConfigCustomSpec", string(data)}, " ")
}
