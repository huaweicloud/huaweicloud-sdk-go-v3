package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationMasksRequest Request Object
type ListNotificationMasksRequest struct {

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	Body *ListNotificationMaskRequestBody `json:"body,omitempty"`
}

func (o ListNotificationMasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMasksRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationMasksRequest", string(data)}, " ")
}
