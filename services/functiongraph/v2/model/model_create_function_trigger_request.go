package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFunctionTriggerRequest Request Object
type CreateFunctionTriggerRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	Body *CreateFunctionTriggerRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionTriggerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionTriggerRequest struct{}"
	}

	return strings.Join([]string{"CreateFunctionTriggerRequest", string(data)}, " ")
}
