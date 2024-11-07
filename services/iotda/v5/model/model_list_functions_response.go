package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionsResponse Response Object
type ListFunctionsResponse struct {

	// **参数说明**：编解码函数列表。
	ProductFunctions *[]FunctionDto `json:"product_functions,omitempty"`

	// **参数说明**：满足查询条件的记录总数。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFunctionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionsResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionsResponse", string(data)}, " ")
}
