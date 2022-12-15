package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type HandleNotificationResponse struct {

	// 请求成功的结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o HandleNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleNotificationResponse struct{}"
	}

	return strings.Join([]string{"HandleNotificationResponse", string(data)}, " ")
}
