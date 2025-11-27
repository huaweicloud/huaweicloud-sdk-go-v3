package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterStatus struct {

	// kubernetes版本
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	// conditions信息
	Conditions *[]ConditionStatus `json:"conditions,omitempty"`

	NodeSummary *NodeSummary `json:"nodeSummary,omitempty"`

	ResourceSummary *ResourceSummary `json:"resourceSummary,omitempty"`

	// 端点
	Endpoints *interface{} `json:"endpoints,omitempty"`

	// 阶段状态信息
	Phase *string `json:"phase,omitempty"`

	// 表述上次状况变化的原因
	Reason *string `json:"reason,omitempty"`

	// 上次状态转换的详细信息
	Message *string `json:"message,omitempty"`

	// 欠费冻结
	ArrearFreeze *string `json:"arrearFreeze,omitempty"`

	// 公安冻结
	PoliceFreeze *string `json:"policeFreeze,omitempty"`

	// 开启的资源列表
	ApiEnablements *[]ApiEnablement `json:"apiEnablements,omitempty"`
}

func (o ClusterStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterStatus struct{}"
	}

	return strings.Join([]string{"ClusterStatus", string(data)}, " ")
}
