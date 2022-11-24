package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteRecordSetsRequest struct {
	Body *BatchDeleteRecordSetsReq `json:"body,omitempty"`
}

func (o BatchDeleteRecordSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetsRequest", string(data)}, " ")
}
