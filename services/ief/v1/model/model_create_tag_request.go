package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTagRequest struct {
	// 资源id

	ResourceId string `json:"resource_id"`
	// 资源类型

	ResourceType string `json:"resource_type"`
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *AddTag `json:"body,omitempty"`
}

func (o CreateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTagRequest", string(data)}, " ")
}
