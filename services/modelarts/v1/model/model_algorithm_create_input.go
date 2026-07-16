package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmCreateInput 算法的数据输入。
type AlgorithmCreateInput struct {

	// 数据输入通道名称。
	Name *string `json:"name,omitempty"`

	// 数据输入通道描述信息。
	Description *string `json:"description,omitempty"`

	// 数据输入约束。
	RemoteConstraints *[]RemoteConstraint `json:"remote_constraints,omitempty"`
}

func (o AlgorithmCreateInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmCreateInput struct{}"
	}

	return strings.Join([]string{"AlgorithmCreateInput", string(data)}, " ")
}
