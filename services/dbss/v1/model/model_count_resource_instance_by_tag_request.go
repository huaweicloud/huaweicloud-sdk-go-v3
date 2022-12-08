package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CountResourceInstanceByTagRequest struct {

	// 资源类型。审计：auditInstance
	ResourceType string `json:"resource_type"`

	Body *ResourceInstanceTagRequest `json:"body,omitempty"`
}

func (o CountResourceInstanceByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourceInstanceByTagRequest struct{}"
	}

	return strings.Join([]string{"CountResourceInstanceByTagRequest", string(data)}, " ")
}
