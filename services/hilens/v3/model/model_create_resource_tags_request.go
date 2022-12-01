package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResourceTagsRequest struct {

	// 资源ID，不同资源（节点，部署，配置项，密钥）有不同的资源ID
	ResourceId string `json:"resource_id"`

	// 资源类型（节点，部署，配置项，密钥）
	ResourceType string `json:"resource_type"`

	Body *TagRequestDetail `json:"body,omitempty"`
}

func (o CreateResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceTagsRequest", string(data)}, " ")
}
