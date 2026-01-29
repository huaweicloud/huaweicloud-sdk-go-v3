package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataclassTypeResponse Response Object
type CreateDataclassTypeResponse struct {

	// 类型id
	Id *string `json:"id,omitempty"`

	// 数据类id
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 账号id
	DomainId *string `json:"domain_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 局点id
	RegionId *string `json:"region_id,omitempty"`

	// 布局id
	LayoutId *string `json:"layout_id,omitempty"`

	// 布局名称
	LayoutName *string `json:"layout_name,omitempty"`

	// 数据类类型分类
	Category *string `json:"category,omitempty"`

	// 数据类类型分类编码
	CategoryCode *string `json:"category_code,omitempty"`

	// 数据类类型名称
	SubCategory *string `json:"sub_category,omitempty"`

	// 数据类类型业务编码
	SubCategoryCode *string `json:"sub_category_code,omitempty"`

	// 数据类类型描述
	Description *string `json:"description,omitempty"`

	// 状态
	Enabled *bool `json:"enabled,omitempty"`

	// 事件等级
	Level *int32 `json:"level,omitempty"`

	// 是否内置，true内置，false非内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 响应时长
	Sla *int64 `json:"sla,omitempty"`

	// 创建人id
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改人id
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改人名称
	ModifierName *string `json:"modifier_name,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 所属数据类业务编码
	DataclassBusinessCode *string `json:"dataclass_business_code,omitempty"`

	// 类型分类下子类型数目
	SubCount       *int32 `json:"sub_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateDataclassTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataclassTypeResponse struct{}"
	}

	return strings.Join([]string{"CreateDataclassTypeResponse", string(data)}, " ")
}
