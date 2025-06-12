package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueExecuteTaskVo 请求的返回的数据对象
type ResultValueExecuteTaskVo struct {

	// 起始记录数 大于 实际总条数时， 值为0， 分页请求才有此值
	Total *int32 `json:"total,omitempty"`

	Value *ExecuteTaskVo `json:"value,omitempty"`

	// 业务失败的提示内容
	Reason *string `json:"reason,omitempty"`

	// 每页展示条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 是否有更多
	HasMore *bool `json:"has_more,omitempty"`
}

func (o ResultValueExecuteTaskVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueExecuteTaskVo struct{}"
	}

	return strings.Join([]string{"ResultValueExecuteTaskVo", string(data)}, " ")
}
