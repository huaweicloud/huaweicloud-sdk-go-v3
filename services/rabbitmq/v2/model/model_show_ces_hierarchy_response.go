package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCesHierarchyResponse Response Object
type ShowCesHierarchyResponse struct {

	// **参数解释**： 监控维度。
	Dimensions *[]ShowCeshierarchyRespDimensions `json:"dimensions,omitempty"`

	// **参数解释**： 实例信息。
	InstanceIds *[]ShowCeshierarchyRespInstanceIds `json:"instance_ids,omitempty"`

	// **参数解释**： 节点信息。
	Nodes *[]ShowCeshierarchyRespNodes `json:"nodes,omitempty"`

	// **参数解释**： Queue信息。
	Queues *[]ShowCeshierarchyRespQueues `json:"queues,omitempty"`

	// **参数解释**： Vhost信息
	Vhosts *[]ShowCeshierarchyRespVhosts `json:"vhosts,omitempty"`

	// **参数解释**： Exchange信息
	Exchanges *[]ShowCeshierarchyRespExchanges `json:"exchanges,omitempty"`

	// **参数解释**： 消费组信息。
	Groups         *[]ShowCeshierarchyRespGroups `json:"groups,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowCesHierarchyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCesHierarchyResponse struct{}"
	}

	return strings.Join([]string{"ShowCesHierarchyResponse", string(data)}, " ")
}
