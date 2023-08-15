package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotificationResponse Response Object
type DeleteNotificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationResponse struct{}"
	}

	return strings.Join([]string{"DeleteNotificationResponse", string(data)}, " ")
}
