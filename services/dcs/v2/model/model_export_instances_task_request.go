package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportInstancesTaskRequest Request Object
type ExportInstancesTaskRequest struct {
	Body *ExportInstancesTaskBody `json:"body,omitempty"`
}

func (o ExportInstancesTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstancesTaskRequest struct{}"
	}

	return strings.Join([]string{"ExportInstancesTaskRequest", string(data)}, " ")
}
