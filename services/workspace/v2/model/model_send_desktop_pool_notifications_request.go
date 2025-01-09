package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendDesktopPoolNotificationsRequest Request Object
type SendDesktopPoolNotificationsRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *SendDesktopPoolNotificationsReq `json:"body,omitempty"`
}

func (o SendDesktopPoolNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDesktopPoolNotificationsRequest struct{}"
	}

	return strings.Join([]string{"SendDesktopPoolNotificationsRequest", string(data)}, " ")
}
