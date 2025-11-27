package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConditionStatus type为Ready，表示集群是否可用；type为Cluster，表示集群的网络连接状态；type为Federation，表示集群的联邦状态。
type ConditionStatus struct {

	// 状态类型
	Type *string `json:"type,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 状态对象的版本
	Observedgeneration *int32 `json:"observedgeneration,omitempty"`

	// 上次状态更新的时间
	LastTransitionTime *sdktime.SdkTime `json:"lastTransitionTime,omitempty"`

	// 状态原因
	Reason *string `json:"reason,omitempty"`

	// 状态信息
	Message *string `json:"message,omitempty"`
}

func (o ConditionStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionStatus struct{}"
	}

	return strings.Join([]string{"ConditionStatus", string(data)}, " ")
}
