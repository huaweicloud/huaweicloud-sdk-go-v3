package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicableInstancesInfo 可被参数模板应用的实例信息
type ApplicableInstancesInfo struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 实例名称。 **取值范围：** 不涉及。
	InstanceName string `json:"instance_name"`

	// **参数解释：** 节点组信息或节点信息的列表对象。 **取值范围：** 当参数模板是副本集或单节点类型时，不返回该字段，模板直接应用到对应实例。 当参数模板是集群类型时，如果是shard组或者config组的参数模板，则返回的是对应类型的节点组信息；如果是mongos组的参数模板，则返回的是对应类型的的节点信息。 例如：一个mongos参数模板可应用到一个或多个mongos节点。
	Entities []EntityInfo `json:"entities"`
}

func (o ApplicableInstancesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicableInstancesInfo struct{}"
	}

	return strings.Join([]string{"ApplicableInstancesInfo", string(data)}, " ")
}
