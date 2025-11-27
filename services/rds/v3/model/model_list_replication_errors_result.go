package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReplicationErrorsResult 发布订阅报错详情。
type ListReplicationErrorsResult struct {

	// 报错发生时间。
	OccurTime *string `json:"occur_time,omitempty"`

	// 错误源类型id。
	SourceTypeId *string `json:"source_type_id,omitempty"`

	// 错误源名称。
	SourceName *string `json:"source_name,omitempty"`

	// 错误代码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误消息。
	ErrorText *string `json:"error_text,omitempty"`
}

func (o ListReplicationErrorsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationErrorsResult struct{}"
	}

	return strings.Join([]string{"ListReplicationErrorsResult", string(data)}, " ")
}
