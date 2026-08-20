package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskIp struct {

	// Task 名称，如 worker-0。
	Task *string `json:"task,omitempty"`

	// Task IP 地址。
	Ip *string `json:"ip,omitempty"`

	// 宿主机 IP。 **约束限制**：仅专属资源池作业返回；公共资源池作业该字段为空。
	HostIp *string `json:"host_ip,omitempty"`

	// 当前 Task 的第几次调度，默认 1。 重调度、抢占等场景下递增。
	ScheduleCount *int32 `json:"schedule_count,omitempty"`
}

func (o TaskIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskIp struct{}"
	}

	return strings.Join([]string{"TaskIp", string(data)}, " ")
}
