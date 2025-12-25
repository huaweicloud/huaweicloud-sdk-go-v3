package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourceInstanceRequest Request Object
type CountResourceInstanceRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-type"`

	Body *QueryResourceInstanceRequestBody `json:"body,omitempty"`
}

func (o CountResourceInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourceInstanceRequest struct{}"
	}

	return strings.Join([]string{"CountResourceInstanceRequest", string(data)}, " ")
}
