package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePublicRecordsetsTaskRequest Request Object
type BatchCreatePublicRecordsetsTaskRequest struct {
	Body *BatchCreatePublicRecordsetsTaskRequestBody `json:"body,omitempty"`
}

func (o BatchCreatePublicRecordsetsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicRecordsetsTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicRecordsetsTaskRequest", string(data)}, " ")
}
