package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotificationSettingResponse Response Object
type CreateNotificationSettingResponse struct {

	// 消息通知配置的唯一标识符。
	Id *string `json:"id,omitempty"`

	// 消息通知配置的唯一资源标识符。
	Urn            *string `json:"urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNotificationSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationSettingResponse struct{}"
	}

	return strings.Join([]string{"CreateNotificationSettingResponse", string(data)}, " ")
}
