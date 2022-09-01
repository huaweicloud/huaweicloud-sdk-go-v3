package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportFunctionRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn" xml:"function_urn"`

	// 是否导出函数配置
	Config *bool `json:"config,omitempty" xml:"config"`

	// 是否导出函数代码
	Code *bool `json:"code,omitempty" xml:"code"`

	// 兼容老的方式，type=code代表导出代码,type=config代码导出配置
	Type *string `json:"type,omitempty" xml:"type"`
}

func (o ExportFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFunctionRequest struct{}"
	}

	return strings.Join([]string{"ExportFunctionRequest", string(data)}, " ")
}
