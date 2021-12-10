package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTasksRequestBody struct {
	// 作业的名称

	Name string `json:"name"`
	// 作业的描述

	Description *string `json:"description,omitempty"`

	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`
	// 作业对应服务的版本号

	ServiceVersion string `json:"service_version"`
	// 仅边缘作业需填且必填，作业运行指定的边缘运行池ID

	EdgePoolId *string `json:"edge_pool_id,omitempty"`
	// 作业指定的算法能力包包周期订单ID

	ResourceOrderId *string `json:"resource_order_id,omitempty"`

	ServiceConfig *TaskServiceConfig `json:"service_config,omitempty"`
}

func (o CreateTasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTasksRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTasksRequestBody", string(data)}, " ")
}
