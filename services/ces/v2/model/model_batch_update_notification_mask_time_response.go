package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateNotificationMaskTimeResponse Response Object
type BatchUpdateNotificationMaskTimeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateNotificationMaskTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNotificationMaskTimeResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateNotificationMaskTimeResponse", string(data)}, " ")
}
