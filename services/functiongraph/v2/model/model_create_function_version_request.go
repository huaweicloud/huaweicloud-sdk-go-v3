package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateFunctionVersionRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	Body *CreateFunctionVersionRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateFunctionVersionRequest", string(data)}, " ")
}
