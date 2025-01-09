package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScreenRecordsRequest Request Object
type BatchDeleteScreenRecordsRequest struct {
	Body *BatchDeleteScreenRecordsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteScreenRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScreenRecordsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteScreenRecordsRequest", string(data)}, " ")
}
