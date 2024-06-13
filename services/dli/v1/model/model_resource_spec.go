package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSpec 表示对作业使用 CPU 和内存资源的包装，依据本数据结构可以生成对 JM/TM 资源的粗粒度的描述。
type ResourceSpec struct {

	// CPU 核数
	Cpu *float64 `json:"cpu,omitempty"`

	// 内存，单位 GiB。
	Memory *string `json:"memory,omitempty"`
}

func (o ResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpec struct{}"
	}

	return strings.Join([]string{"ResourceSpec", string(data)}, " ")
}
