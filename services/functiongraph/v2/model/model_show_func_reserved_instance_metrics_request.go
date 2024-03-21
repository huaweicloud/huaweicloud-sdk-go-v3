package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFuncReservedInstanceMetricsRequest Request Object
type ShowFuncReservedInstanceMetricsRequest struct {

	// 本次查询起始位置，默认值0
	Marker *string `json:"marker,omitempty"`

	// 本次查询最大返回的数据条数，最大值500，默认值100
	Limit *string `json:"limit,omitempty"`

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FuncUrn string `json:"func_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowFuncReservedInstanceMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFuncReservedInstanceMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowFuncReservedInstanceMetricsRequest", string(data)}, " ")
}
