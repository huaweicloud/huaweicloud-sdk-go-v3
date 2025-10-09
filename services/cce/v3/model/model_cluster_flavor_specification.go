package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterFlavorSpecification **参数解释**： 集群可售卖规格详情
type ClusterFlavorSpecification struct {

	// **参数解释**： 规格名称 **取值范围**： - cce.s1.small: 小规模单控制节点CCE集群（最大50节点） - cce.s1.medium: 中等规模单控制节点CCE集群（最大200节点） - cce.s2.small: 小规模三控制节点CCE集群（最大50节点） - cce.s2.medium: 中等规模三控制节点CCE集群（最大200节点） - cce.s2.large: 大规模三控制节点CCE集群（最大1000节点） - cce.s2.xlarge: 超大规模三控制节点CCE集群（最大2000节点） [- cce.s3.small: 小规模五控制节点CCE集群（最大50节点）](tag:hcs,hcs_sm) [- cce.s3.medium: 中等规模五控制节点CCE集群（最大200节点）](tag:hcs,hcs_sm) [- cce.s3.large: 大规模五控制节点CCE集群（最大1000节点）](tag:hcs,hcs_sm) [- cce.s3.xlarge: 超大规模五控制节点CCE集群（最大2000节点）](tag:hcs,hcs_sm)  [专属云特殊规格如下：](tag:hws,hws_hk,hcs,hcs_sm) [- cce.dec.s1.small: 小规模单控制节点的专属云CCE集群（最大50节点）](tag:hws,hws_hk) [- cce.dec.s1.medium: 中等规模单控制节点的专属云CCE集群（最大200节点）](tag:hws,hws_hk) [- cce.dec.s1.large: 大规模单控制节点的专属云CCE集群（最大1000节点）](tag:hws,hws_hk) [- cce.dec.s1.xlarge: 超大规模单控制节点的专属云CCE集群（最大2000节点）](tag:hws,hws_hk) [- cce.dec.s2.small: 小规模三控制节点的专属云CCE集群（最大50节点）](tag:hws,hws_hk) [- cce.dec.s2.medium: 中等规模三控制节点的专属云CCE集群（最大200节点）](tag:hws,hws_hk) [- cce.dec.s2.large: 大规模三控制节点的专属云CCE集群（最大1000节点）](tag:hws,hws_hk) [- cce.dec.s2.xlarge: 超大规模三控制节点的专属云CCE集群（最大2000节点）](tag:hws,hws_hk) [- cce.dec.s3.small: 小规模五控制节点的专属云CCE集群（最大50节点）](tag:hcs,hcs_sm) [- cce.dec.s3.medium: 中等规模五控制节点的专属云CCE集群（最大200节点）](tag:hcs,hcs_sm) [- cce.dec.s3.large: 大规模五控制节点的专属云CCE集群（最大1000节点）](tag:hcs,hcs_sm) [- cce.dec.s3.xlarge: 超大规模五控制节点的专属云CCE集群（最大2000节点）](tag:hcs,hcs_sm)
	Name *string `json:"name,omitempty"`

	// **参数解释**： 集群节点规模 **取值范围**： - 50: 最大支持50节点 - 200: 最大支持200节点 - 100: 最大支持1000节点 - 2000: 最大支持2000节点
	NodeCapacity *int32 `json:"nodeCapacity,omitempty"`

	// **参数解释**： 控制节点详情
	AvailableMasterFlavors *[]MasterFlavorSpec `json:"availableMasterFlavors,omitempty"`

	// **参数解释**： 集群规格是否售罄 **取值范围**： 不涉及
	IsSoldOut *bool `json:"isSoldOut,omitempty"`

	// **参数解释**： 是否支持控制节点多可用区分布 **取值范围**： 不涉及
	IsSupportMultiAZ *bool `json:"isSupportMultiAZ,omitempty"`
}

func (o ClusterFlavorSpecification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterFlavorSpecification struct{}"
	}

	return strings.Join([]string{"ClusterFlavorSpecification", string(data)}, " ")
}
