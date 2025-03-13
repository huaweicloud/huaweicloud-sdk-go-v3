package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradePlanSpec struct {

	// 集群ID
	ClusterID string `json:"clusterID"`

	// 当前集群版本
	ClusterVersion string `json:"clusterVersion"`

	// 本次集群升级的目标版本
	TargetVersion string `json:"targetVersion"`

	// 自动升级计划的最早时间（UTC时间），需要早于notBeforeDeadline
	NotBefore string `json:"notBefore"`

	// 自动升级计划的最晚时间（UTC时间）
	NotAfter string `json:"notAfter"`

	// 自动升级计划开始的截止时间（UTC时间）
	NotBeforeDeadline string `json:"notBeforeDeadline"`
}

func (o UpgradePlanSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradePlanSpec struct{}"
	}

	return strings.Join([]string{"UpgradePlanSpec", string(data)}, " ")
}
