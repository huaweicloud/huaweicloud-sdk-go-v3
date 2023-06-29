package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlossaryListRequest Request Object
type ShowGlossaryListRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 标签类型 缺省值：all
	Type *string `json:"type,omitempty"`

	// 标签名
	Name *string `json:"name,omitempty"`

	// 标签创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 开始时间
	Start *string `json:"start,omitempty"`

	// 结束时间
	End *string `json:"end,omitempty"`

	// 分页参数:每页限定数量 缺省值：10
	Limit *string `json:"limit,omitempty"`

	// 分页参数：页数 缺省值：0
	Offset *string `json:"offset,omitempty"`

	// 排序字段 默认为createTime 缺省值：createTime
	SortBy *string `json:"sort_by,omitempty"`

	// 排序方式 默认排序字段为降序 缺省值：desc
	SortOrder *string `json:"sort_order,omitempty"`
}

func (o ShowGlossaryListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlossaryListRequest struct{}"
	}

	return strings.Join([]string{"ShowGlossaryListRequest", string(data)}, " ")
}
