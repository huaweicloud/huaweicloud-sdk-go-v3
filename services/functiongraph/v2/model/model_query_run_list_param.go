package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRunListParam struct {

	// 页码
	Page int32 `json:"page"`

	// 每页大小
	PageSize int32 `json:"page_size"`

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
