package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListOperateRecordRequestBody struct {

	// 查询开始时间，格式为毫秒级时间戳。
	StartTime int64 `json:"start_time"`

	// 查询结束时间，格式为毫秒级时间戳。
	EndTime int64 `json:"end_time"`

	// 操作类型
	OperateType *string `json:"operate_type,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 事件等级
	Level *string `json:"level,omitempty"`

	// 查询偏移量，默认为0。
	Offset *string `json:"offset,omitempty"`

	// 查询数量，默认为10。
	Limit *string `json:"limit,omitempty"`

	// 排序字段，默认为operate_time。取值范围：operate_type、user_name、operate_time、level。
	Sort *string `json:"sort,omitempty"`

	// 排序方式，默认为desc（倒序）。取值范围：desc（倒序）、asc（正序）。
	Order *string `json:"order,omitempty"`
}

func (o ListOperateRecordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperateRecordRequestBody struct{}"
	}

	return strings.Join([]string{"ListOperateRecordRequestBody", string(data)}, " ")
}
