package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderAlertEventContent 事件内容
type OrderAlertEventContent struct {

	// 状态
	HandleStatus *string `json:"handle_status,omitempty"`
}

func (o OrderAlertEventContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderAlertEventContent struct{}"
	}

	return strings.Join([]string{"OrderAlertEventContent", string(data)}, " ")
}
