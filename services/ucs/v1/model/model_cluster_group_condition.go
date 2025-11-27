package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterGroupCondition 舰队启用联邦或权限策略信息
type ClusterGroupCondition struct {

	// 类型。 - Federation：舰队启用联邦状态类型 - Policy：权限策略
	Type *string `json:"type,omitempty"`

	// 舰队启用的联邦或权限策略的状态
	Status *string `json:"status,omitempty"`

	// 状态原因
	Reason *string `json:"reason,omitempty"`

	// 状态信息
	Message *string `json:"message,omitempty"`

	// 状态更新时间
	LastTransitionTime *sdktime.SdkTime `json:"lastTransitionTime,omitempty"`
}

func (o ClusterGroupCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterGroupCondition struct{}"
	}

	return strings.Join([]string{"ClusterGroupCondition", string(data)}, " ")
}
