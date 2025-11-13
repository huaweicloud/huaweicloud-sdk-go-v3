package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePublicRecordsetsTaskRequest Request Object
type BatchUpdatePublicRecordsetsTaskRequest struct {
	Body *BatchUpdatePublicRecordsetsTaskRequestBody `json:"body,omitempty"`
}

func (o BatchUpdatePublicRecordsetsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePublicRecordsetsTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdatePublicRecordsetsTaskRequest", string(data)}, " ")
}
