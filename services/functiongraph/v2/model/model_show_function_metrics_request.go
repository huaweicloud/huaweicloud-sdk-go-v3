package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFunctionMetricsRequest Request Object
type ShowFunctionMetricsRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FuncUrn string `json:"func_urn"`

	// 时间间隔（单位：min）
	Period string `json:"period"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowFunctionMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowFunctionMetricsRequest", string(data)}, " ")
}
