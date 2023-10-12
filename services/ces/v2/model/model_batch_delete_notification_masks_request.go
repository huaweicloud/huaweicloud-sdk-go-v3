package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteNotificationMasksRequest Request Object
type BatchDeleteNotificationMasksRequest struct {
	Body *BatchDeleteNotificationMasksRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteNotificationMasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteNotificationMasksRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteNotificationMasksRequest", string(data)}, " ")
}
