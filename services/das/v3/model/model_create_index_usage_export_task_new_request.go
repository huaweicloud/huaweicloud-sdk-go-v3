package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndexUsageExportTaskNewRequest Request Object
type CreateIndexUsageExportTaskNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateIndexUsageExportTaskNewRequestBody `json:"body,omitempty"`
}

func (o CreateIndexUsageExportTaskNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndexUsageExportTaskNewRequest struct{}"
	}

	return strings.Join([]string{"CreateIndexUsageExportTaskNewRequest", string(data)}, " ")
}
