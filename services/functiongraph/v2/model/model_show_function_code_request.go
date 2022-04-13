package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowFunctionCodeRequest struct {
	// 函数的URN，详细解释见FunctionGraph函数模型的描述。

	FunctionUrn string `json:"function_urn"`
}

func (o ShowFunctionCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionCodeRequest struct{}"
	}

	return strings.Join([]string{"ShowFunctionCodeRequest", string(data)}, " ")
}
