package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddDeleteTagsRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
	// 资源类型 - ief-edge_node - ief-deployment - ief-application - ief-device

	ResourceType string `json:"resource_type"`
	// 资源ID

	ResourceId string `json:"resource_id"`

	Body *BachTags `json:"body,omitempty"`
}

func (o BatchAddDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddDeleteTagsRequest", string(data)}, " ")
}
