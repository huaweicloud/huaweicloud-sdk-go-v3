package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点对象。
type ClusterListInstances struct {
	// 状态。  - 100：创建中。 - 200：可用。 - 303：不可用，如创建失败。

	Status *string `json:"status,omitempty"`
	// 支持类型：ess（Elasticsearch节点）。

	Type *string `json:"type,omitempty"`
	// 实例ID。

	Id *string `json:"id,omitempty"`
	// 实例名字。

	Name *string `json:"name,omitempty"`
	// 节点规格名称。

	SpecCode *string `json:"specCode,omitempty"`
	// 节点所属AZ信息。

	AzCode *string `json:"azCode,omitempty"`
}

func (o ClusterListInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterListInstances struct{}"
	}

	return strings.Join([]string{"ClusterListInstances", string(data)}, " ")
}
