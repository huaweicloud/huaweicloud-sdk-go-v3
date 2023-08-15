package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyConfigurationRequestBody struct {

	// 实例ID或组ID或节点ID的列表对象。可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 - 当变更的实例类型是集群，如果变更的是shard组或者config组的参数模板，传值为组ID。如果变更的是mongos节点的参数模板，传值为节点ID。 - 当变更的实例类型是副本集或单节点，传值为实例ID。
	EntityIds []string `json:"entity_ids"`
}

func (o ApplyConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationRequestBody", string(data)}, " ")
}
