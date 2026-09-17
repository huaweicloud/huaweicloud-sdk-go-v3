package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlLimitingRecordInfo SqlLimitingRecordInfo对象
type SqlLimitingRecordInfo struct {

	// SQL限流规则ID
	ItemId *string `json:"item_id,omitempty"`

	// SQL类型
	Type *string `json:"type,omitempty"`

	// 限流规则
	KeyStr *string `json:"key_str,omitempty"`

	// 最大并发数
	MaxConnection *string `json:"max_connection,omitempty"`

	// 最大等待时间
	MaxWaiting *string `json:"max_waiting,omitempty"`

	// 当前并发数
	CurConnection *int32 `json:"cur_connection,omitempty"`

	// 当前拦截数
	CurReject *int32 `json:"cur_reject,omitempty"`

	// 总拦截数
	TotalReject *int32 `json:"total_reject,omitempty"`

	// 创建时间
	CreateAt *string `json:"create_at,omitempty"`

	// PostgreSQL限流语句标准化后唯一标识
	QueryId *string `json:"query_id,omitempty"`
}

func (o SqlLimitingRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlLimitingRecordInfo struct{}"
	}

	return strings.Join([]string{"SqlLimitingRecordInfo", string(data)}, " ")
}
