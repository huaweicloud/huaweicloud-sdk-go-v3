package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlgorithmResponseJobConfigOutputs struct {

	// 数据输出通道名称。
	Name string `json:"name"`

	// 数据输出通道描述信息。
	Description *string `json:"description,omitempty"`
}

func (o AlgorithmResponseJobConfigOutputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmResponseJobConfigOutputs struct{}"
	}

	return strings.Join([]string{"AlgorithmResponseJobConfigOutputs", string(data)}, " ")
}
