package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTagRequest Request Object
type BatchDeleteResourceTagRequest struct {

	// 资源类型。审计：auditInstance
	ResourceType string `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	Body *ResourceTagRequest `json:"body,omitempty"`
}

func (o BatchDeleteResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagRequest", string(data)}, " ")
}
