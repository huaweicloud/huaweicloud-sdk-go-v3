package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestoreRecordsRequest Request Object
type ListRestoreRecordsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 查询开始时间。格式：yyyyMMddHHmmss，如：20170718235959。
	BeginTime *string `json:"begin_time,omitempty"`

	// 查询结束时间。格式：yyyyMMddHHmmss，如：20170718235959。
	EndTime *string `json:"end_time,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRestoreRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListRestoreRecordsRequest", string(data)}, " ")
}
