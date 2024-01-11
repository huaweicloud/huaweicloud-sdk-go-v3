package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueUserReq 资源队列用户。
type WorkloadQueueUserReq struct {

	// 工作队列名称。
	QueueName string `json:"queue_name"`

	// 资源队列用户列表
	UserList []WorkloadQueueUserReqUserList `json:"user_list"`
}

func (o WorkloadQueueUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueUserReq struct{}"
	}

	return strings.Join([]string{"WorkloadQueueUserReq", string(data)}, " ")
}
