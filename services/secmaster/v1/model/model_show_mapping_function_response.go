package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMappingFunctionResponse Response Object
type ShowMappingFunctionResponse struct {

	// 比较函数信息
	CompareList *[]DpeCompareFunctionDetail `json:"compare_list,omitempty"`

	// 操作函数信息
	OperationList  *[]DpeOperateFunctionDetail `json:"operation_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowMappingFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMappingFunctionResponse struct{}"
	}

	return strings.Join([]string{"ShowMappingFunctionResponse", string(data)}, " ")
}
