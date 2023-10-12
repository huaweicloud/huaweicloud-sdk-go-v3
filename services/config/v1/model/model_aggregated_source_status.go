package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregatedSourceStatus 资源聚合器状态响应体。
type AggregatedSourceStatus struct {

	// 源帐号最近一次聚合失败时返回的错误码。
	LastErrorCode *string `json:"last_error_code,omitempty"`

	// 源帐号最近一次聚合失败时返回的错误消息。
	LastErrorMessage *string `json:"last_error_message,omitempty"`

	// 最近一次更新的状态类型。
	LastUpdateStatus *string `json:"last_update_status,omitempty"`

	// 最近一次更新的时间。
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 源帐号ID或组织。
	SourceId *string `json:"source_id,omitempty"`

	// 账号名。
	SourceName *string `json:"source_name,omitempty"`

	// 源帐号类型（ACCOUNT | ORGANIZATION）。
	SourceType *string `json:"source_type,omitempty"`
}

func (o AggregatedSourceStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregatedSourceStatus struct{}"
	}

	return strings.Join([]string{"AggregatedSourceStatus", string(data)}, " ")
}
