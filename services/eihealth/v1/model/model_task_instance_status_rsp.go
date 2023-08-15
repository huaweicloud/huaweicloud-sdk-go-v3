package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceStatusRsp struct {

	// 实例执行状态
	Phase *string `json:"phase,omitempty"`

	// 实例IP
	PodIp *string `json:"pod_ip,omitempty"`

	// 实例所在节点IP
	HostIp *string `json:"host_ip,omitempty"`

	// 计算节点的名称
	HostName *string `json:"host_name,omitempty"`

	// 实例创建时间
	StartTime *string `json:"start_time,omitempty"`

	// 实例状态信息
	ContainerStatuses *[]TaskInstanceContainerStatusRsp `json:"container_statuses,omitempty"`
}

func (o TaskInstanceStatusRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceStatusRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceStatusRsp", string(data)}, " ")
}
