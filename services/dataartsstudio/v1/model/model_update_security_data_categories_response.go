package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityDataCategoriesResponse Response Object
type UpdateSecurityDataCategoriesResponse struct {

	// 数据分类id。
	CategoryId *string `json:"category_id,omitempty"`

	// 数据分类名称。
	CategoryName *string `json:"category_name,omitempty"`

	// 数据分类描述。
	Description *string `json:"description,omitempty"`

	// 数据分类层级。
	CategoryLevel *int32 `json:"category_level,omitempty"`

	// 分类树根节点。
	RootId *string `json:"root_id,omitempty"`

	// 父分类节点。
	ParentId *string `json:"parent_id,omitempty"`

	// 分类树路径。
	CategoryPath *string `json:"category_path,omitempty"`

	// 创建者。
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新者。
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间。
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateSecurityDataCategoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityDataCategoriesResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityDataCategoriesResponse", string(data)}, " ")
}
