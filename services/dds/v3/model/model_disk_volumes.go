package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskVolumes struct {

	// 实例ID或组ID或节点ID。可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 - 当变更的实例类型是集群，如果变更的是shard组或者config组的参数模板，传值为组ID。如果变更的是mongos节点的参数模板，传值为节点ID。 - 当变更的实例类型是副本集或单节点，传值为实例ID。
	EntityId string `json:"entity_id"`

	// 实例名称或组名称或节点名称
	EntityName string `json:"entity_name"`

	// group_type。取值范围： - mongos，表示集群mongos节点类型。 - shard，表示集群shard节点类型。 - config，表示集群config节点类型。 - replica，表示副本集类型。 - single，表示单节点类型。 - readonly，表示只读节点类型。
	GroupType string `json:"group_type"`

	// 使用量，保留两位小数，单位(GB)
	Used string `json:"used"`

	// 总大小，单位(GB)
	Size string `json:"size"`
}

func (o DiskVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskVolumes struct{}"
	}

	return strings.Join([]string{"DiskVolumes", string(data)}, " ")
}
