package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayFederationRayClusterRequest **参数解释**：Ray联邦端点中RayCluster端点的信息（请求）。 **约束限制**：不涉及。
type RayFederationRayClusterRequest struct {

	// **参数解释**：RayCluster端点ID，可通过查询Ray联邦端点列表接口获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文、数字、中划线的组合。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o RayFederationRayClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayFederationRayClusterRequest struct{}"
	}

	return strings.Join([]string{"RayFederationRayClusterRequest", string(data)}, " ")
}
