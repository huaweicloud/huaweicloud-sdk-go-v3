package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFlinkJobsRequest Request Object
type ImportFlinkJobsRequest struct {
	Body *ImportFlinkJobRequestBody `json:"body,omitempty"`
}

func (o ImportFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"ImportFlinkJobsRequest", string(data)}, " ")
}
