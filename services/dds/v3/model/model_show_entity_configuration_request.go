package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEntityConfigurationRequest Request Object
type ShowEntityConfigurationRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type"`

	// 实例ID，可以调用“查询实例列表和详情-QueryingInstancesandDetails”接口获取。如果未申请实例，可以调用“创建实例-CreatingaDBInstance”接口创建。
	InstanceId string `json:"instance_id"`

	// - 实例ID或组ID或节点ID。可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 - 当获取的实例类型是集群，如果获取的是shard组或者config组的参数模板，传值为组ID。如果获取的是mongos节点的参数模板，传值为节点ID。 - 当获取的实例类型是副本集或单节点，传值为实例ID。
	EntityId string `json:"entity_id"`
}

func (o ShowEntityConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEntityConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowEntityConfigurationRequest", string(data)}, " ")
}
