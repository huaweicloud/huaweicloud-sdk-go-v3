package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCesHierarchyResponse struct {

	// 监控维度。
	Dimensions *[]ShowCeshierarchyRespDimensions `json:"dimensions,omitempty" xml:"dimensions"`

	// 实例信息。
	InstanceIds *[]ShowCeshierarchyRespInstanceIds `json:"instance_ids,omitempty" xml:"instance_ids"`

	// 节点信息。
	Nodes *[]ShowCeshierarchyRespNodes `json:"nodes,omitempty" xml:"nodes"`

	// 队列信息。
	Queues *[]ShowCeshierarchyRespQueues `json:"queues,omitempty" xml:"queues"`

	// 消费组信息。
	Groups         *[]ShowCeshierarchyRespGroups `json:"groups,omitempty" xml:"groups"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowCesHierarchyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCesHierarchyResponse struct{}"
	}

	return strings.Join([]string{"ShowCesHierarchyResponse", string(data)}, " ")
}
