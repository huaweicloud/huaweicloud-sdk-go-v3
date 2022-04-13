package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTasksRequest struct {
	// 服务名称

	ServiceName string `json:"service_name"`

	Body *CreateTasksRequestBody `json:"body,omitempty"`
}

func (o CreateTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTasksRequest struct{}"
	}

	return strings.Join([]string{"CreateTasksRequest", string(data)}, " ")
}
