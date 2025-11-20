package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotificationSettingResponse Response Object
type DeleteNotificationSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNotificationSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationSettingResponse struct{}"
	}

	return strings.Join([]string{"DeleteNotificationSettingResponse", string(data)}, " ")
}
