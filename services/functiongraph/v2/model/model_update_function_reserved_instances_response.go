package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateFunctionReservedInstancesResponse struct {
	// 预留实例个数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateFunctionReservedInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionReservedInstancesResponse struct{}"
	}

	return strings.Join([]string{"UpdateFunctionReservedInstancesResponse", string(data)}, " ")
}
