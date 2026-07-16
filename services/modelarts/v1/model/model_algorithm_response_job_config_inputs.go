package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlgorithmResponseJobConfigInputs struct {

	// 数据输入通道名称。
	Name string `json:"name"`

	// 数据输入通道描述信息。
	Description *string `json:"description,omitempty"`

	// 数据输入约束。
	RemoteConstraints *[]AlgorithmResponseJobConfigRemoteConstraints `json:"remote_constraints,omitempty"`
}

func (o AlgorithmResponseJobConfigInputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmResponseJobConfigInputs struct{}"
	}

	return strings.Join([]string{"AlgorithmResponseJobConfigInputs", string(data)}, " ")
}
