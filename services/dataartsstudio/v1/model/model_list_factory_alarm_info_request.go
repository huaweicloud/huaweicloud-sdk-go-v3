package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryAlarmInfoRequest Request Object
type ListFactoryAlarmInfoRequest struct {

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 告警的开始时间，默认当前时间的前一个小时，13位时间戳
	StartTime *int64 `json:"start_time,omitempty"`

	// 告警的最后时间，默认为当前时间，13位时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 分页的起始页，默认值为0。取值范围大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页返回结果，指定每页最大记录数。默认值10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFactoryAlarmInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryAlarmInfoRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryAlarmInfoRequest", string(data)}, " ")
}
