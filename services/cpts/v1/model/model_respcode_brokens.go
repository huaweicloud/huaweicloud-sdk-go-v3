package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RespcodeBrokens struct {

	// 校验失败
	CheckPointFailed *[]float64 `json:"checkPointFailed,omitempty" xml:"checkPointFailed"`

	// 异常请求
	Error *[]float64 `json:"error,omitempty" xml:"error"`

	// 其他失败
	OthersFailed *[]float64 `json:"othersFailed,omitempty" xml:"othersFailed"`

	// 解析失败
	ParsedFailed *[]float64 `json:"parsedFailed,omitempty" xml:"parsedFailed"`

	// 连接被拒
	RefusedFailed *[]float64 `json:"refusedFailed,omitempty" xml:"refusedFailed"`

	// 成功请求
	Success *[]float64 `json:"success,omitempty" xml:"success"`

	// 超时失败
	Timeout *[]float64 `json:"timeout,omitempty" xml:"timeout"`
}

func (o RespcodeBrokens) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespcodeBrokens struct{}"
	}

	return strings.Join([]string{"RespcodeBrokens", string(data)}, " ")
}
