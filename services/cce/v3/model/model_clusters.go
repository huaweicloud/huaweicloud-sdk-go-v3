package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Clusters struct {

	// **参数解释**： 集群名字。 **约束限制**： 不涉及 **取值范围**： - internalCluster：私网访问的集群 - externalCluster：公网访问的集群  **默认取值**： - 若不存在publicIp（虚拟机弹性IP），则集群列表的集群数量为1，该字段值为“internalCluster”。 - 若存在publicIp，则集群列表的集群数量大于1，所有扩展的cluster的name的值为“externalCluster”。
	Name *string `json:"name,omitempty"`

	Cluster *ClusterCert `json:"cluster,omitempty"`
}

func (o Clusters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Clusters struct{}"
	}

	return strings.Join([]string{"Clusters", string(data)}, " ")
}
