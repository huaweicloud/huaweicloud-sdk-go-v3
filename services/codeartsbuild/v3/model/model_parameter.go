package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Parameter 构建作业参数
type Parameter struct {

	// 参数定义名，默认为hudson.model.StringParameterDefinition
	Name *string `json:"name,omitempty"`

	// 构建执行参数子参数
	Params *[]CreateBuildJobParameterParam `json:"params,omitempty"`
}

func (o Parameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Parameter struct{}"
	}

	return strings.Join([]string{"Parameter", string(data)}, " ")
}
