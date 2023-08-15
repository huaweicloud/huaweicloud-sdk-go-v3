package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeFunctionRequest Request Object
type InvokeFunctionRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 取值为：tail（返回函数执行后的4K日志），或者为空（不返回日志）。
	XCffLogType *string `json:"X-Cff-Log-Type,omitempty"`

	// 返回体格式，取值v0,v1。 v0:默认返回文本格式 v1:默认返回json格式，sdk需要使用此值。
	XCFFRequestVersion *string `json:"X-CFF-Request-Version,omitempty"`

	// 执行函数请求体，为json格式。
	Body map[string]interface{} `json:"body,omitempty"`
}

func (o InvokeFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeFunctionRequest struct{}"
	}

	return strings.Join([]string{"InvokeFunctionRequest", string(data)}, " ")
}
