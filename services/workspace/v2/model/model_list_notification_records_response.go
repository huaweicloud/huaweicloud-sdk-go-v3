package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationRecordsResponse Response Object
type ListNotificationRecordsResponse struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 通知拦截记录信息
	NotificationRecords *[]NotificationRecordInfo `json:"notification_records,omitempty"`
	HttpStatusCode      int                       `json:"-"`
}

func (o ListNotificationRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationRecordsResponse", string(data)}, " ")
}
