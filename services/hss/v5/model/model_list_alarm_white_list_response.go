package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmWhiteListResponse Response Object
type ListAlarmWhiteListResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 可继续添加的白名单数量
	RemainNum *int32 `json:"remain_num,omitempty"`

	// 白名单数量上限
	LimitNum *int32 `json:"limit_num,omitempty"`

	// 支持筛选的事件类型
	EventTypeList *[]int32 `json:"event_type_list,omitempty"`

	// 告警白名单详情
	DataList       *[]AlarmWhiteListResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListAlarmWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmWhiteListResponse", string(data)}, " ")
}
