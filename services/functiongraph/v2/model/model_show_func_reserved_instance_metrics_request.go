package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFuncReservedInstanceMetricsRequest Request Object
type ShowFuncReservedInstanceMetricsRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FuncUrn string `json:"func_urn"`
}

func (o ShowFuncReservedInstanceMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFuncReservedInstanceMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowFuncReservedInstanceMetricsRequest", string(data)}, " ")
}
