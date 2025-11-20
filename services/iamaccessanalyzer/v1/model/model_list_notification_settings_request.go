package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationSettingsRequest Request Object
type ListNotificationSettingsRequest struct {

	// 单页最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListNotificationSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationSettingsRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationSettingsRequest", string(data)}, " ")
}
