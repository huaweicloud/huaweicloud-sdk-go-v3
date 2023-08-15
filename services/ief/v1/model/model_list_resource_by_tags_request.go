package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceByTagsRequest Request Object
type ListResourceByTagsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 资源类型 - ief-edge_node - ief-deployment - ief-application - ief-device
	ResourceType string `json:"resource_type"`

	Body *Tags `json:"body,omitempty"`
}

func (o ListResourceByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceByTagsRequest", string(data)}, " ")
}
