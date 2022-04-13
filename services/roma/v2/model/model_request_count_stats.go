package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RequestCountStats struct {
	// 请求总次数

	ReqCount *int32 `json:"req_count,omitempty"`
	// 2xx响应码总次数

	ReqCount2xx *int32 `json:"req_count2xx,omitempty"`
	// 4xx响应码总次数

	ReqCount4xx *int32 `json:"req_count4xx,omitempty"`
	// 5xx响应码总次数

	ReqCount5xx *int32 `json:"req_count5xx,omitempty"`
	// 错误次数

	ReqCountError *int32 `json:"req_count_error,omitempty"`
}

func (o RequestCountStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestCountStats struct{}"
	}

	return strings.Join([]string{"RequestCountStats", string(data)}, " ")
}
