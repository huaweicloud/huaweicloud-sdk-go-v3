package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstanceByTagRequest Request Object
type ListResourceInstanceByTagRequest struct {

	// 资源类型。审计：auditInstance
	ResourceType string `json:"resource_type"`

	// 查询记录数（action为count时无此参数）如果action为filter默认为1000，limit最多为1000,不能为负数，最小值为1
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量（action为count时无此参数）从第一条数据偏移offset条数据后开始查询，如果action为filter默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数
	Offset *string `json:"offset,omitempty"`

	Body *ResourceInstanceTagRequest `json:"body,omitempty"`
}

func (o ListResourceInstanceByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstanceByTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstanceByTagRequest", string(data)}, " ")
}
