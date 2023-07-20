package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckConnectionRequest Request Object
type CheckConnectionRequest struct {

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *VerityConnectivityReq `json:"body,omitempty"`
}

func (o CheckConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckConnectionRequest struct{}"
	}

	return strings.Join([]string{"CheckConnectionRequest", string(data)}, " ")
}
