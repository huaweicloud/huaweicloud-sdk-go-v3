package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 改变探针状态的参数。
type AgentStatusChangeParam struct {

	// 探针实例id列表。
	InstanceList []int64 `json:"instance_list"`

	// 期望探针改变后的状态，0或1，0表示启用，1表示停用。
	TargetStatus int32 `json:"target_status"`

	// 探针所在的区域。
	Region string `json:"region"`

	// 探针所属环境的id。
	EnvId *int64 `json:"env_id,omitempty"`
}

func (o AgentStatusChangeParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentStatusChangeParam struct{}"
	}

	return strings.Join([]string{"AgentStatusChangeParam", string(data)}, " ")
}
