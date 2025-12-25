package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTagValueRequest Request Object
type UpdateTagValueRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 标签键
	Key string `json:"key"`

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	Body *UpdateTagValueRequestBody `json:"body,omitempty"`
}

func (o UpdateTagValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTagValueRequest struct{}"
	}

	return strings.Join([]string{"UpdateTagValueRequest", string(data)}, " ")
}
