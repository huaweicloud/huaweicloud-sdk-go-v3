package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFunctionReservedInstancesRequestBody struct {
	// 预留实例个数

	Count int32 `json:"count"`
}

func (o UpdateFunctionReservedInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionReservedInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionReservedInstancesRequestBody", string(data)}, " ")
}
