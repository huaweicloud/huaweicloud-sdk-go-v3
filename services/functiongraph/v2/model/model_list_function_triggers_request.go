package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFunctionTriggersRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn" xml:"function_urn"`
}

func (o ListFunctionTriggersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionTriggersRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionTriggersRequest", string(data)}, " ")
}
