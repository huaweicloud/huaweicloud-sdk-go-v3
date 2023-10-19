package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagRequest struct {

	// 标签id
	TagIds *[]string `json:"tag_ids,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 标签名称，用作搜索框筛选
	Name *string `json:"name,omitempty"`

	// 创建者，用作搜索框筛选
	CreateUser *string `json:"create_user,omitempty"`

	// 开始时间
	Start *string `json:"start,omitempty"`

	// 结束时间
	End *string `json:"end,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 每页大小
	Limit *int32 `json:"limit,omitempty"`

	// 根据xx排序
	SortBy *string `json:"sort_by,omitempty"`

	// 升序/降序
	SortOrder *string `json:"sort_order,omitempty"`
}

func (o TagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagRequest struct{}"
	}

	return strings.Join([]string{"TagRequest", string(data)}, " ")
}
