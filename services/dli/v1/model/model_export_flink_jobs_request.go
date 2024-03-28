package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFlinkJobsRequest Request Object
type ExportFlinkJobsRequest struct {
	Body *ExportFlinkJobsRequestBody `json:"body,omitempty"`
}

func (o ExportFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"ExportFlinkJobsRequest", string(data)}, " ")
}
