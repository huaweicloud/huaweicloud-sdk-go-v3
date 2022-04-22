package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListNotificationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`
}

func (o ListNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationRequest", string(data)}, " ")
}
