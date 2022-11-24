package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 可被参数模板应用的实例信息
type ApplicableInstancesInfo struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 节点组信息或节点信息的列表对象。  当参数模板是集群类型时，如果是shard组或者config组的参数模板，则可应用到的是对应类型的节点组，如果是mongos组的参数模板，则可应用到的是对应类型的的节点。  当参数模板是副本集或单节点类型时，直接应用到对应实例。  例如：一个mongos参数模板可应用到一个或多个mongos节点。
	Entities []EntityInfo `json:"entities"`
}

func (o ApplicableInstancesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicableInstancesInfo struct{}"
	}

	return strings.Join([]string{"ApplicableInstancesInfo", string(data)}, " ")
}
