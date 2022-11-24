package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListReservedInstanceConfigsRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn *string `json:"function_urn,omitempty"`
}

func (o ListReservedInstanceConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReservedInstanceConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListReservedInstanceConfigsRequest", string(data)}, " ")
}
