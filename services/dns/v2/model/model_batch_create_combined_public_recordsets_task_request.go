package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateCombinedPublicRecordsetsTaskRequest Request Object
type BatchCreateCombinedPublicRecordsetsTaskRequest struct {
	Body *BatchCreateCombinedPublicRecordsetsTaskRequestBody `json:"body,omitempty"`
}

func (o BatchCreateCombinedPublicRecordsetsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCombinedPublicRecordsetsTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateCombinedPublicRecordsetsTaskRequest", string(data)}, " ")
}
