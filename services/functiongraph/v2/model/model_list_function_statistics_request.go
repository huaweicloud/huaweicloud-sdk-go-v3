package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFunctionStatisticsRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FuncUrn string `json:"func_urn"`
	// 获取最近多少分钟内函数执行的指标。

	Period string `json:"period"`
}

func (o ListFunctionStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionStatisticsRequest", string(data)}, " ")
}
