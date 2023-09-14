package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFunctionUrlRequest Request Object
type CreateFunctionUrlRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	Body *FunctionUrlRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionUrlRequest struct{}"
	}

	return strings.Join([]string{"CreateFunctionUrlRequest", string(data)}, " ")
}
