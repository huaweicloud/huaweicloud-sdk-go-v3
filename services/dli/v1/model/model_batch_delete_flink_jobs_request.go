package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFlinkJobsRequest Request Object
type BatchDeleteFlinkJobsRequest struct {
	Body *BatchDeleteFlinkJobsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteFlinkJobsRequest", string(data)}, " ")
}
