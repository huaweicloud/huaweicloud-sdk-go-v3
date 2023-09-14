package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationRecordsRequest Request Object
type ListNotificationRecordsRequest struct {
}

func (o ListNotificationRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationRecordsRequest", string(data)}, " ")
}
