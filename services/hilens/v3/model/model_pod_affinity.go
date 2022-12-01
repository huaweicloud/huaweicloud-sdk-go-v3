package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PodAffinity struct {

	// 节点亲和规则
	NodeAffinity *[]MatchExpression `json:"node_affinity,omitempty"`

	// Pod亲和规则
	PodAffinity *[]MatchExpression `json:"pod_affinity,omitempty"`
}

func (o PodAffinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodAffinity struct{}"
	}

	return strings.Join([]string{"PodAffinity", string(data)}, " ")
}
