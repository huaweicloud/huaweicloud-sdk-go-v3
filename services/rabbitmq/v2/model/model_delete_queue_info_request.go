package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueueInfoRequest Request Object
type DeleteQueueInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// Vhost名称
	Vhost string `json:"vhost"`

	// Queue名称
	Queue string `json:"queue"`
}

func (o DeleteQueueInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueInfoRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueueInfoRequest", string(data)}, " ")
}
