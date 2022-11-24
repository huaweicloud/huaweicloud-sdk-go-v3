package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFunctionReservedInstancesResponse struct {

	// 函数预留实例列表
	Reservedinstances *[]FuncReservedInstance `json:"reservedinstances,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 函数个数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFunctionReservedInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionReservedInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionReservedInstancesResponse", string(data)}, " ")
}
