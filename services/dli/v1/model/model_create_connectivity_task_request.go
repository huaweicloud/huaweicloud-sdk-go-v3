package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectivityTaskRequest Request Object
type CreateConnectivityTaskRequest struct {

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *CreateConnectivityTaskRequestBody `json:"body,omitempty"`
}

func (o CreateConnectivityTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectivityTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectivityTaskRequest", string(data)}, " ")
}
