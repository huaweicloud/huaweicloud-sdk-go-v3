package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConfigurationParameterResult struct {

	// 实例ID或组ID或节点ID。可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 - 当变更的实例类型是集群，如果变更的是shard组或者config组的参数模板，传值为组ID。如果变更的是mongos节点的参数模板，传值为节点ID。 - 当变更的实例类型是副本集或单节点，传值为实例ID。
	EntityId string `json:"entity_id" xml:"entity_id"`

	// 参数名和参数值映射关系。用户可以基于默认参数模板的参数，自定义的参数值。
	ParameterValues map[string]string `json:"parameter_values" xml:"parameter_values"`
}

func (o UpdateConfigurationParameterResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterResult struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterResult", string(data)}, " ")
}
