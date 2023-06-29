package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationTemplatesResponse Response Object
type ListNotificationTemplatesResponse struct {

	// 模板数组
	Results        *[]NotificationTemplate `json:"results,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListNotificationTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationTemplatesResponse", string(data)}, " ")
}
