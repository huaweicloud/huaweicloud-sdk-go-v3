package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBuildJobParameter 创建构建作业参数
type UpdateBuildJobParameter struct {

	// 参数定义名，默认为hudson.model.StringParameterDefinition
	Name *string `json:"name,omitempty"`

	// 构建执行参数子参数
	Params *[]UpdateBuildJobParameterParam `json:"params,omitempty"`
}

func (o UpdateBuildJobParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBuildJobParameter struct{}"
	}

	return strings.Join([]string{"UpdateBuildJobParameter", string(data)}, " ")
}
