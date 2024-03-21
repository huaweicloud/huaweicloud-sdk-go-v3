package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionStatisticsRequest Request Object
type ListFunctionStatisticsRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FuncUrn string `json:"func_urn"`

	// 获取最近多少分钟内函数执行的指标。
	Period string `json:"period"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListFunctionStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionStatisticsRequest", string(data)}, " ")
}
