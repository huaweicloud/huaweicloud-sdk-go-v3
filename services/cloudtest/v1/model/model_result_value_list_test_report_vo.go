package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueListTestReportVo 请求的返回的数据对象
type ResultValueListTestReportVo struct {

	// 起始记录数 大于 实际总条数时， 值为0， 分页请求才有此值
	Total *int32 `json:"total,omitempty"`

	// 实际的数据类型：单个对象，集合 或 NULL
	Value *[]TestReportVo `json:"value,omitempty"`

	// 业务失败的提示内容，对内接口才有此值
	Reason *string `json:"reason,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNo *int32 `json:"page_no,omitempty"`

	HasMore *bool `json:"has_more,omitempty"`
}

func (o ResultValueListTestReportVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueListTestReportVo struct{}"
	}

	return strings.Join([]string{"ResultValueListTestReportVo", string(data)}, " ")
}
