package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Workload 工作负载信息
type Workload struct {

	// 工作负载ID
	WorkloadId string `json:"workload_id"`

	// 工作负载名称
	WorkloadName string `json:"workload_name"`

	// 工作负载类型,包含如下：   - deployments：无状态负载   - statefulsets：有状态负载   - daemonsets：守护进程表
	WorkloadType string `json:"workload_type"`
}

func (o Workload) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workload struct{}"
	}

	return strings.Join([]string{"Workload", string(data)}, " ")
}
