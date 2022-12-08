package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddResourceTagRequest struct {

	// 资源类型。审计：auditInstance
	ResourceType string `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	Body *ResourceTagRequest `json:"body,omitempty"`
}

func (o BatchAddResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddResourceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchAddResourceTagRequest", string(data)}, " ")
}
