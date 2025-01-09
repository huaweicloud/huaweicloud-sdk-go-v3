package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendNotificationsResponse Response Object
type SendNotificationsResponse struct {

	// 发送桌面消息任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SendNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendNotificationsResponse struct{}"
	}

	return strings.Join([]string{"SendNotificationsResponse", string(data)}, " ")
}
