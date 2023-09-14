package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionUrlRequest Request Object
type DeleteFunctionUrlRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`
}

func (o DeleteFunctionUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionUrlRequest struct{}"
	}

	return strings.Join([]string{"DeleteFunctionUrlRequest", string(data)}, " ")
}
