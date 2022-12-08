package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRunListParam struct {

	// 偏移量,表示从此偏移量开始查询,默认值为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页大小
	Limit int32 `json:"limit"`

	// 查询开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o QueryRunListParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRunListParam struct{}"
	}

	return strings.Join([]string{"QueryRunListParam", string(data)}, " ")
}
