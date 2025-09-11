package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditSummaryStatusRequest struct {

	// 操作码，operType为switch时必输入  - on: 开启  - off: 关闭
	Code *string `json:"code,omitempty"`

	// 操作类型  - switch: 任务开关  - execute: 立即刷新
	OperType string `json:"oper_type"`

	// 查询条件：开始时间
	QueryBeginTime *int64 `json:"query_begin_time,omitempty"`

	// 查询条件：结束时间
	QueryEndTime *int64 `json:"query_end_time,omitempty"`
}

func (o AuditSummaryStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditSummaryStatusRequest struct{}"
	}

	return strings.Join([]string{"AuditSummaryStatusRequest", string(data)}, " ")
}
