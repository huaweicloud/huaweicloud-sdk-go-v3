package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerNetworkHyperCluster **参数解释**：创建服务器的参数面网络信息。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type ServerNetworkHyperCluster struct {

	// 参数解释：HyperCluster的id。 约束限制：不涉及。 取值范围：不涉及。 默认取值：不涉及。
	Id *string `json:"id,omitempty"`

	// 参数解释：HyperCluster的子网id。 约束限制：不涉及。 取值范围：不涉及。 默认取值：不涉及。
	HyperClusterSubnetId *string `json:"hyper_cluster_subnet_id,omitempty"`
}

func (o ServerNetworkHyperCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerNetworkHyperCluster struct{}"
	}

	return strings.Join([]string{"ServerNetworkHyperCluster", string(data)}, " ")
}
