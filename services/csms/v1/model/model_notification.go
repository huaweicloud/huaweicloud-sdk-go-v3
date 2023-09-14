package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Notification 通知主题对象。
type Notification struct {

	// 事件通知的对象类型。
	TargetType string `json:"target_type"`

	// 事件通知的对象ID。
	TargetId string `json:"target_id"`

	// 事件通知的对象名称。
	TargetName string `json:"target_name"`
}

func (o Notification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Notification struct{}"
	}

	return strings.Join([]string{"Notification", string(data)}, " ")
}
