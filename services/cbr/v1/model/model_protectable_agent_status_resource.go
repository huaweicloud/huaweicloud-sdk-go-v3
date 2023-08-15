package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectableAgentStatusResource struct {

	// 待检查资源ID
	ResourceId string `json:"resource_id"`

	// 待检查资源name
	ResourceName *string `json:"resource_name,omitempty"`

	// 待检查的资源类型。当前支持的取值包含两个：OS::Nova::Server，该值代表保护的资源为云服务器，OS::Ironic::BareMetalServer，该值代表保护的资源为裸金属服务器。
	ResourceType string `json:"resource_type"`
}

func (o ProtectableAgentStatusResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableAgentStatusResource struct{}"
	}

	return strings.Join([]string{"ProtectableAgentStatusResource", string(data)}, " ")
}
