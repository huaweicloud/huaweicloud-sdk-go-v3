package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFlavorByTypeReq struct {

	// 变更后节点规格ID。 该参数通过[获取实例规格列表](ListFlavors.xml)接口获取，根据name属性所需要的规格，选择对应的flavor_id。 仅支持同一个Esasticsearch引擎版本下的节点规格变更。
	NewFlavorId string `json:"new_flavor_id"`

	// 操作类型。
	OperationType *string `json:"operation_type,omitempty"`

	// 是否需要检查副本，取值范围为true或false。默认开启校验。 - true: 开启副本校验。 - false: 忽略副本校验。  >Master和Client节点不是数据节点，因此不需要进行副本校验。即使选择true，也不会进行副本校验。
	NeedCheckReplica *bool `json:"need_check_replica,omitempty"`

	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`

	// 是否需要检查集群状态，取值范围为true或false。默认开启校验。
	NeedCheckClusterStatus *bool `json:"need_check_cluster_status,omitempty"`

	// 是否需要检查集群负载，取值范围为true或false。默认开启校验。
	ClusterLoadCheck *bool `json:"cluster_load_check,omitempty"`
}

func (o UpdateFlavorByTypeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlavorByTypeReq struct{}"
	}

	return strings.Join([]string{"UpdateFlavorByTypeReq", string(data)}, " ")
}
