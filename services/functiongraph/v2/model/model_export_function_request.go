package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFunctionRequest Request Object
type ExportFunctionRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 是否导出函数配置，默认为false
	Config *bool `json:"config,omitempty"`

	// 是否导出函数代码，默认为false
	Code *bool `json:"code,omitempty"`

	// 不兼容与code、config参数混用；type=code代表导出代码,type=config代码导出配置
	Type *string `json:"type,omitempty"`
}

func (o ExportFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFunctionRequest struct{}"
	}

	return strings.Join([]string{"ExportFunctionRequest", string(data)}, " ")
}
