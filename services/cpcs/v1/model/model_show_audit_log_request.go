package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditLogRequest Request Object
type ShowAuditLogRequest struct {

	// 指定查询返回记录条数，默认值10
	PageSize *int32 `json:"page_size,omitempty"`

	// 索引位置，从page_num指定的下一条数据开始查询默认值为0
	PageNum *int32 `json:"page_num,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ShowAuditLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditLogRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditLogRequest", string(data)}, " ")
}
