package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationRecordsRequest Request Object
type ListNotificationRecordsRequest struct {

	// 每页返回的个数。  默认值：50。
	Limit *string `json:"limit,omitempty"`

	// 分页查询起始的事件通知记录时间，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`
}

func (o ListNotificationRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationRecordsRequest", string(data)}, " ")
}
