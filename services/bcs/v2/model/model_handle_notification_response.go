package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type HandleNotificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HandleNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleNotificationResponse struct{}"
	}

	return strings.Join([]string{"HandleNotificationResponse", string(data)}, " ")
}
