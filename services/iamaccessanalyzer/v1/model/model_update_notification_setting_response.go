package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationSettingResponse Response Object
type UpdateNotificationSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNotificationSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationSettingResponse", string(data)}, " ")
}
