package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFunctionConfigRequest struct {
	// 函数的URN，详细解释见FunctionGraph函数模型。

	FunctionUrn string `json:"function_urn"`

	Body *UpdateFunctionConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionConfigRequest", string(data)}, " ")
}
