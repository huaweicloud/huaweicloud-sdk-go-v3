package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordIndexRequest Request Object
type CreateRecordIndexRequest struct {
	Body *RecordIndexRequestBody `json:"body,omitempty"`
}

func (o CreateRecordIndexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordIndexRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordIndexRequest", string(data)}, " ")
}
