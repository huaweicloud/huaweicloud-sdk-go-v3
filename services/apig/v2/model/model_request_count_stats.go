package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RequestCountStats struct {

	// 请求总次数
	ReqCount *int32 `json:"req_count,omitempty" xml:"req_count"`

	// 2xx响应码总次数
	ReqCount2xx *int32 `json:"req_count2xx,omitempty" xml:"req_count2xx"`

	// 4xx响应码总次数
	ReqCount4xx *int32 `json:"req_count4xx,omitempty" xml:"req_count4xx"`

	// 5xx响应码总次数
	ReqCount5xx *int32 `json:"req_count5xx,omitempty" xml:"req_count5xx"`

	// 错误次数
	ReqCountError *int32 `json:"req_count_error,omitempty" xml:"req_count_error"`
}

func (o RequestCountStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestCountStats struct{}"
	}

	return strings.Join([]string{"RequestCountStats", string(data)}, " ")
}
