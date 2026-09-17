package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMissingIndexExportTaskNewRequest Request Object
type CreateMissingIndexExportTaskNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateMissingIndexExportTaskNewRequestBody `json:"body,omitempty"`
}

func (o CreateMissingIndexExportTaskNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMissingIndexExportTaskNewRequest struct{}"
	}

	return strings.Join([]string{"CreateMissingIndexExportTaskNewRequest", string(data)}, " ")
}
