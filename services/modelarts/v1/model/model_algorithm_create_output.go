package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmCreateOutput 算法的数据输出。
type AlgorithmCreateOutput struct {

	// 数据输出通道名称。
	Name string `json:"name"`

	// 数据输出通道描述信息。
	Description *string `json:"description,omitempty"`
}

func (o AlgorithmCreateOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmCreateOutput struct{}"
	}

	return strings.Join([]string{"AlgorithmCreateOutput", string(data)}, " ")
}
