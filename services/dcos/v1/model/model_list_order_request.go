package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrderRequest Request Object
type ListOrderRequest struct {

	// 服务类型
	Type *string `json:"type,omitempty"`

	// 服务条目
	SubType *string `json:"sub_type,omitempty"`

	// 阶段
	Stage *string `json:"stage,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 申请人
	Applicant *string `json:"applicant,omitempty"`

	// 时间范围-开始
	StartTime *string `json:"start_time,omitempty"`

	// 时间范围-结束
	EndTime *string `json:"end_time,omitempty"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 单页大小
	Limit *int32 `json:"limit,omitempty"`

	// 横移量
	Offset *int32 `json:"offset,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrderRequest struct{}"
	}

	return strings.Join([]string{"ListOrderRequest", string(data)}, " ")
}
