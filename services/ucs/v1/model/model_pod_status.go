package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PodStatus 用于表示 Pod 状态信息
type PodStatus struct {

	// 表示Pod的生命周期状态，有 5 种取值： - Pending：Pod 已被接受但尚未完全就绪 - Running：Pod 已绑定到节点且至少有一个容器在运行 - Succeeded：Pod 中所有容器都已成功终止，且不会被重启 - Failed：Pod 中至少一个容器失败终止 - Unknown：无法获取 Pod 状态
	Phase *string `json:"phase,omitempty"`

	// Pod所在主机的IP地址
	HostIP *string `json:"hostIP,omitempty"`

	// 记录Pod被Kubelet认可的时间
	StartTime *string `json:"startTime,omitempty"`

	// 容器的状态列表
	ContainerStatuses *[]ContainerStatus `json:"containerStatuses,omitempty"`
}

func (o PodStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodStatus struct{}"
	}

	return strings.Join([]string{"PodStatus", string(data)}, " ")
}
