package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferDeploymentPodsRequest Request Object
type ListInferDeploymentPodsRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 部署名称，在创建部署时即可在返回体中获取，也可通过[查询服务部署列表](ListInferDeployments.xml)获取当前用户拥有的部署，其name字段即为部署名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DeploymentName string `json:"deployment_name"`

	// **参数解释：** 服务实例名字，可以为all，为all时去查询所有的服务实例。 **约束限制：** 不涉及。 **取值范围：** 服务实例名字。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** pod状态，一次支持多种状态筛选，多种状态以\",\"连接，不能存在空格。默认不过滤。取值范围有7种RUNNING（运行中）、PENDING（未就绪）、SUCCEEDED（成功）、FAILED（失败）、ABNORMAL（异常）、UNKNOWN（未知）、DELETED（已删除）。 **约束限制：** 不涉及。
	Status *[]string `json:"status,omitempty"`

	// **参数解释：** 指定每一页返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表的起始页。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** pod名字。 **取值范围：** 不涉及。
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释：** pod ID。 **取值范围：** 不涉及。
	PodId *string `json:"pod_id,omitempty"`

	// **参数解释：** pod节点IP地址。 **取值范围：** 不涉及。
	PodNodeIp *string `json:"pod_node_ip,omitempty"`

	// **参数解释：** pod节点名称。 **取值范围：** 不涉及。
	PodNodeName *string `json:"pod_node_name,omitempty"`
}

func (o ListInferDeploymentPodsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferDeploymentPodsRequest struct{}"
	}

	return strings.Join([]string{"ListInferDeploymentPodsRequest", string(data)}, " ")
}
