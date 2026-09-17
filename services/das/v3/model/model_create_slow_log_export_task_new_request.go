package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSlowLogExportTaskNewRequest Request Object
type CreateSlowLogExportTaskNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateSlowLogExportTaskNewRequestBody `json:"body,omitempty"`
}

func (o CreateSlowLogExportTaskNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSlowLogExportTaskNewRequest struct{}"
	}

	return strings.Join([]string{"CreateSlowLogExportTaskNewRequest", string(data)}, " ")
}
