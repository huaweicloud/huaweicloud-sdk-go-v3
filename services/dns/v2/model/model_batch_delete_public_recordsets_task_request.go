package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePublicRecordsetsTaskRequest Request Object
type BatchDeletePublicRecordsetsTaskRequest struct {
	Body *BatchDeletePublicRecordsetsTaskRequestBody `json:"body,omitempty"`
}

func (o BatchDeletePublicRecordsetsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicRecordsetsTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicRecordsetsTaskRequest", string(data)}, " ")
}
