package model

import (
	"encoding/json"

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
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFunctionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionStatisticsRequest", string(data)}, " ")
}
