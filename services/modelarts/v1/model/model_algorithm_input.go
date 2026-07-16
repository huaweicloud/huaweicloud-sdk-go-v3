package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmInput 算法输入通道信息。
type AlgorithmInput struct {

	// 数据输入通道名称。
	Name string `json:"name"`

	// 数据输入输出通道映射的容器本地路径。
	LocalDir *string `json:"local_dir,omitempty"`

	Remote *AlgorithmRemote `json:"remote"`
}

func (o AlgorithmInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmInput struct{}"
	}

	return strings.Join([]string{"AlgorithmInput", string(data)}, " ")
}
