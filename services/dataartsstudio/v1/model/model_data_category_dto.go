package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataCategoryDto struct {

	// 分类id
	CategoryId *string `json:"category_id,omitempty"`

	// 分类名称
	CategoryName *string `json:"category_name,omitempty"`

	// 分类描述
	Description *string `json:"description,omitempty"`

	// 分类层级
	CategoryLevel *int32 `json:"category_level,omitempty"`

	// 识别规则
	RuleList *[]DataClassificationRuleQueryDto `json:"rule_list,omitempty"`

	// 分类树根节点
	RootId *string `json:"root_id,omitempty"`

	// 父分类节点
	ParentId *string `json:"parent_id,omitempty"`

	// 分类树路径
	CategoryPath *string `json:"category_path,omitempty"`

	// 创建者
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新者
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 同步（默认都同步资产）
	Synchronize *bool `json:"synchronize,omitempty"`

	// 子分类
	Children *[]DataCategoryDto `json:"children,omitempty"`
}

func (o DataCategoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataCategoryDto struct{}"
	}

	return strings.Join([]string{"DataCategoryDto", string(data)}, " ")
}
