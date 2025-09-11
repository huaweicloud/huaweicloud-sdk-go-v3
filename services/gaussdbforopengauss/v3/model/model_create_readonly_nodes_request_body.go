package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateReadonlyNodesRequestBody struct {

	// **参数解释**: 新创建只读节点在各可用区的分布情况。 **约束限制**: 不涉及。
	NodeDistribution []NodeDistributionOption `json:"node_distribution"`
}

func (o CreateReadonlyNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReadonlyNodesRequestBody struct{}"
	}

	return strings.Join([]string{"CreateReadonlyNodesRequestBody", string(data)}, " ")
}
