package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeDetectRiskResponseInfo struct {

	// **参数解释**： 有风险镜像数量 **取值范围**： 0~2147483647
	ImagesNum *int32 `json:"images_num,omitempty"`

	// **参数解释**： 基线检查风险项总和 **取值范围**： 0~2147483647
	BaselineRiskNum *int32 `json:"baseline_risk_num,omitempty"`

	// **参数解释**： 漏洞数量 **取值范围**： 0~2147483647
	VulNum *int32 `json:"vul_num,omitempty"`

	// **参数解释**： 告警数量 **取值范围**： 0~2147483647
	EventNum *int32 `json:"event_num,omitempty"`

	// **参数解释**： 集群开启防护节点数量 **取值范围**： 0~2147483647
	ProtectNodeNum *int32 `json:"protect_node_num,omitempty"`

	// **参数解释**： 集群节点总数 **取值范围**： 0~2147483647
	NodeTotalNum *int32 `json:"node_total_num,omitempty"`

	// **参数解释**： 集群ID **取值范围**： 字符长度0-128位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 付费模式 **取值范围**： - on_demand：按需 - free：免费
	ChargingMode *string `json:"charging_mode,omitempty"`
}

func (o NodeDetectRiskResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeDetectRiskResponseInfo struct{}"
	}

	return strings.Join([]string{"NodeDetectRiskResponseInfo", string(data)}, " ")
}
