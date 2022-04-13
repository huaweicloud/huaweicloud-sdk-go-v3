package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteFunctionTriggersRequest struct {
	// 函数的URN，详细解释见FunctionGraph函数模型的描述。

	FunctionUrn string `json:"function_urn"`
}

func (o BatchDeleteFunctionTriggersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFunctionTriggersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteFunctionTriggersRequest", string(data)}, " ")
}
