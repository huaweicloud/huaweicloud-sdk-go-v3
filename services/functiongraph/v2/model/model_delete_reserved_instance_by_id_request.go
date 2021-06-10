package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteReservedInstanceByIdRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`
	// 预留实例id。

	InstanceId string `json:"instance_id"`
}

func (o DeleteReservedInstanceByIdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteReservedInstanceByIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteReservedInstanceByIdRequest", string(data)}, " ")
}
