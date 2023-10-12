package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateNotificationMaskTimeRequest Request Object
type BatchUpdateNotificationMaskTimeRequest struct {
	Body *BatchUpdateNotificationMaskTimeRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateNotificationMaskTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNotificationMaskTimeRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateNotificationMaskTimeRequest", string(data)}, " ")
}
