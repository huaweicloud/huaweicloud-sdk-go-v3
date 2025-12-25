package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTagRequest Request Object
type ListProjectTagRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-type"`
}

func (o ListProjectTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTagRequest", string(data)}, " ")
}
