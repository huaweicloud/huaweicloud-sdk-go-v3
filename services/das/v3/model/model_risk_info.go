package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskInfo RiskInfo对象
type RiskInfo struct {

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	InstanceInfo *InstanceInfoForRisk `json:"instance_info,omitempty"`

	// 指标值
	Values *[]float64 `json:"values,omitempty"`
}

func (o RiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskInfo struct{}"
	}

	return strings.Join([]string{"RiskInfo", string(data)}, " ")
}
