package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PodSpec PodSpec 是 Kubernetes 中描述 Pod 的核心数据结构
type PodSpec struct {

	// 定义 Pod 可以挂载的存储卷列表
	Volumes *[]interface{} `json:"volumes,omitempty"`

	// Pod中的主要容器列表
	Containers *[]interface{} `json:"containers,omitempty"`

	// 容器失败后的重启策略
	RestartPolicy *string `json:"restartPolicy,omitempty"`

	// 容器终止前的优雅退出时间
	TerminationGracePeriodSeconds *int32 `json:"terminationGracePeriodSeconds,omitempty"`

	// Pod在节点上的最大存活时间
	ActiveDeadlineSeconds *int32 `json:"activeDeadlineSeconds,omitempty"`

	// DNS策略
	DnsPolicy *string `json:"dnsPolicy,omitempty"`

	// 节点选择器，用于指定Pod可以调度到哪些节点
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// 运行此Pod所使用的ServiceAccount名称
	ServiceAccountName *string `json:"serviceAccountName,omitempty"`

	// 提供向后兼容的旧版ServiceAccount字段
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// 是否自动挂载ServiceAccount的令牌
	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty"`

	// 指定Pod应该调度到的节点
	NodeName *string `json:"nodeName,omitempty"`

	// Pod级别的安全上下文
	SecurityContext *interface{} `json:"securityContext,omitempty"`

	// 指定使用的调度器
	SchedulerName *string `json:"schedulerName,omitempty"`

	// 容器容忍的污点列表
	Tolerations *[]Toleration `json:"tolerations,omitempty"`

	// Pod 的优先级
	Priority *int32 `json:"priority,omitempty"`

	// 服务信息是否注入到Pod的环境变量中
	EnableServiceLinks *bool `json:"enableServiceLinks,omitempty"`

	// 抢占优先级策略
	PreemptionPolicy *string `json:"preemptionPolicy,omitempty"`
}

func (o PodSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodSpec struct{}"
	}

	return strings.Join([]string{"PodSpec", string(data)}, " ")
}
