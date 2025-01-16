package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetRecordSetsStatusRequest Request Object
type BatchSetRecordSetsStatusRequest struct {
	Body *BatchSetRecordSetsStatusRequestBody `json:"body,omitempty"`
}

func (o BatchSetRecordSetsStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetRecordSetsStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchSetRecordSetsStatusRequest", string(data)}, " ")
}
