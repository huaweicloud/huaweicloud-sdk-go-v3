package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsByResourceTypeRequest Request Object
type ListTagsByResourceTypeRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 资源类型 - ief-edge_node - ief-deployment - ief-application - ief-device
	ResourceType string `json:"resource_type"`
}

func (o ListTagsByResourceTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsByResourceTypeRequest struct{}"
	}

	return strings.Join([]string{"ListTagsByResourceTypeRequest", string(data)}, " ")
}
