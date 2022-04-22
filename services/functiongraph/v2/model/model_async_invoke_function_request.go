package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AsyncInvokeFunctionRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 执行函数请求体，为json格式。
	Body map[string]interface{} `json:"body,omitempty"`
}

func (o AsyncInvokeFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncInvokeFunctionRequest struct{}"
	}

	return strings.Join([]string{"AsyncInvokeFunctionRequest", string(data)}, " ")
}
