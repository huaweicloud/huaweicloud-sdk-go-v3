package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateNotificationMasksRequest Request Object
type BatchUpdateNotificationMasksRequest struct {
	Body *BatchUpdateNotificationMasksRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateNotificationMasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNotificationMasksRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateNotificationMasksRequest", string(data)}, " ")
}
