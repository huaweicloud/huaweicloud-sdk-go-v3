package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceTagsRequest Request Object
type ShowResourceTagsRequest struct {

	// 资源ID，不同资源（节点，部署，配置项，密钥）有不同的资源ID
	ResourceId string `json:"resource_id"`

	// 资源类型（节点，部署，配置项，密钥）
	ResourceType string `json:"resource_type"`
}

func (o ShowResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsRequest", string(data)}, " ")
}
