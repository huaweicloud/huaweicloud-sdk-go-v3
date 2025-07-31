package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeDetectRiskResponseInfo struct {

	// 有风险镜像数量
	ImagesNum *int32 `json:"images_num,omitempty"`

	// 基线检查风险项总和
	BaselineRiskNum *int32 `json:"baseline_risk_num,omitempty"`

	// 漏洞数量
	VulNum *int32 `json:"vul_num,omitempty"`

	// 告警数量
	EventNum *int32 `json:"event_num,omitempty"`

	// 集群开启防护节点数量
	ProtectNodeNum *int32 `json:"protect_node_num,omitempty"`

	// 集群节点总数
	NodeTotalNum *int32 `json:"node_total_num,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 付费模式 | on_demand 按需 free 免费
	ChargingMode *string `json:"charging_mode,omitempty"`
}

func (o NodeDetectRiskResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeDetectRiskResponseInfo struct{}"
	}

	return strings.Join([]string{"NodeDetectRiskResponseInfo", string(data)}, " ")
}
