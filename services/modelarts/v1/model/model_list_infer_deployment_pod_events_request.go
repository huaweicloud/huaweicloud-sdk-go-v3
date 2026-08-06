package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferDeploymentPodEventsRequest Request Object
type ListInferDeploymentPodEventsRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 部署名称，在创建部署时即可在返回体中获取，也可通过查询服务部署列表获取当前用户拥有的部署，其name字段即为部署名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DeploymentName string `json:"deployment_name"`

	// **参数解释：** 服务实例名字，可以为all，为all时去查询所有的服务实例。 **约束限制：** 不涉及。 **取值范围：** 服务实例名字。 **默认取值：** 不涉及。
	InstanceName string `json:"instance_name"`

	// **参数解释：** 服务实例pod名字，可以为all，为all时去查询所有的服务实例。 **约束限制：** 不涉及。 **取值范围：** 服务实例名字。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 查询开始时间，Unix时间戳（毫秒）。 **约束限制：** 需要与end_time同时传入或同时为空。不能早于end_time。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释：** 查询结束时间，Unix时间戳（毫秒）。 **约束限制：** 需要与start_time同时传入或同时为空。不能早于start_time。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListInferDeploymentPodEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferDeploymentPodEventsRequest struct{}"
	}

	return strings.Join([]string{"ListInferDeploymentPodEventsRequest", string(data)}, " ")
}
