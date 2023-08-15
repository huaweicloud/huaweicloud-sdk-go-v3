package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeConnectivityRequest Request Object
type ShowNodeConnectivityRequest struct {
	QueueName string `json:"queue_name"`

	TaskId string `json:"task_id"`
}

func (o ShowNodeConnectivityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeConnectivityRequest struct{}"
	}

	return strings.Join([]string{"ShowNodeConnectivityRequest", string(data)}, " ")
}
