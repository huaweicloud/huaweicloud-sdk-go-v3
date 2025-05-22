package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogicalClusterRequestBody **参数解释**： 创建逻辑集群对象。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type CreateLogicalClusterRequestBody struct {
	LogicalCluster *CreateLogicalClusterInfo `json:"logical_cluster"`
}

func (o CreateLogicalClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogicalClusterRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLogicalClusterRequestBody", string(data)}, " ")
}
