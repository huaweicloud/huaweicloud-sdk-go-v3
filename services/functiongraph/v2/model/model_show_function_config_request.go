package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowFunctionConfigRequest struct {
	// 函数的URN，详细解释见FunctionGraph函数模型的描述。

	FunctionUrn string `json:"function_urn"`
}

func (o ShowFunctionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowFunctionConfigRequest", string(data)}, " ")
}
