package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTagsRequest struct {
	// 资源ID

	ResourceId string `json:"resource_id"`
	// 资源类型 1. ief-edge_node 2. ief-deployment 3. ief-application 4. ief-device

	ResourceType string `json:"resource_type"`
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}
