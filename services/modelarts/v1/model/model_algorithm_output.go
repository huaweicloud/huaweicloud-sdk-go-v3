package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmOutput 算法输出通道信息。
type AlgorithmOutput struct {

	// 数据输出通道名称。
	Name string `json:"name"`

	// 数据输出通道映射的容器本地路径。
	LocalDir *string `json:"local_dir,omitempty"`

	Remote *Remote `json:"remote"`

	// 数据传输模式，默认为“upload_periodically”。
	Mode *string `json:"mode,omitempty"`

	// 数据传输周期，默认为30s。
	Period *string `json:"period,omitempty"`
}

func (o AlgorithmOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmOutput struct{}"
	}

	return strings.Join([]string{"AlgorithmOutput", string(data)}, " ")
}
