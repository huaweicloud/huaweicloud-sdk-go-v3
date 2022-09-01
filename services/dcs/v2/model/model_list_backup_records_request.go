package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBackupRecordsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 查询开始时间，时间为UTC时间。格式：yyyyMMddHHmmss，如：20170718235959。
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 查询结束时间，时间为UTC时间。格式：yyyyMMddHHmmss，如：20170718235959。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListBackupRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupRecordsRequest", string(data)}, " ")
}
