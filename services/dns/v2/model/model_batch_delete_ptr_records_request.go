package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePtrRecordsRequest Request Object
type BatchDeletePtrRecordsRequest struct {
	Body *BatchDeletePtrRecordsRequestBody `json:"body,omitempty"`
}

func (o BatchDeletePtrRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePtrRecordsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePtrRecordsRequest", string(data)}, " ")
}
