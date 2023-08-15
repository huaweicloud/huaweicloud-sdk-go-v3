package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTasksRequestBody struct {

	// 作业的名称，必填。仅能包含汉字、字母、数字、中划线和下划线，长度介于1~100之间。
	Name string `json:"name"`

	// 作业的描述，选填。长度不超过500。
	Description *string `json:"description,omitempty"`

	// 作业对应服务的版本号，必填。由两个介于0~999的整数和一个点号分隔符组成。
	ServiceVersion string `json:"service_version"`

	// 作业运行指定的边缘运行池ID，仅边缘场景需填且必填。
	EdgePoolId *string `json:"edge_pool_id,omitempty"`

	// 作业指定的算法能力包包周期订单ID，仅部分服务需填且必填。
	ResourceOrderId *string `json:"resource_order_id,omitempty"`

	Timing *TaskTiming `json:"timing,omitempty"`

	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	ServiceConfig *TaskServiceConfig `json:"service_config,omitempty"`
}

func (o CreateTasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTasksRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTasksRequestBody", string(data)}, " ")
}
