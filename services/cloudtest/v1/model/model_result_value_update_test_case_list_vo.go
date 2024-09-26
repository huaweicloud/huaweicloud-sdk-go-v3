package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueUpdateTestCaseListVo 请求的返回的数据对象
type ResultValueUpdateTestCaseListVo struct {

	// 起始记录数 大于 实际总条数时， 值为0， 分页请求才有此值
	Total *int32 `json:"total,omitempty"`

	Value *UpdateTestCaseListVo `json:"value,omitempty"`

	// 业务失败的提示内容，对内接口才有此值
	Reason *string `json:"reason,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNo *int32 `json:"page_no,omitempty"`

	HasMore *bool `json:"has_more,omitempty"`
}

func (o ResultValueUpdateTestCaseListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueUpdateTestCaseListVo struct{}"
	}

	return strings.Join([]string{"ResultValueUpdateTestCaseListVo", string(data)}, " ")
}
