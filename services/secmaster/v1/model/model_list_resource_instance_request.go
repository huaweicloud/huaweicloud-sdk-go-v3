package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstanceRequest Request Object
type ListResourceInstanceRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 查询记录数（action为count时无此参数）如果action为filter默认为1000，limit最多为1000,不能为负数，最小值为1
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量（action为count时无此参数）从第一条数据偏移offset条数据后开始查询，如果action为filter默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-type"`

	Body *QueryResourceInstanceRequestBody `json:"body,omitempty"`
}

func (o ListResourceInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstanceRequest", string(data)}, " ")
}
