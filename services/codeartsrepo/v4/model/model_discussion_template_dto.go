package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiscussionTemplateDto 检视意见模板返回体
type DiscussionTemplateDto struct {

	// **参数解释：** 检视意见模板主键id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 模板名称。
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释：** 描述。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 默认指派人。
	AssigneeId *int32 `json:"assignee_id,omitempty"`

	// **参数解释：** 意见分类key。
	Categories *string `json:"categories,omitempty"`

	// **参数解释：** 意见分类英文。
	CategoriesEn *string `json:"categories_en,omitempty"`

	// **参数解释：** 意见分类中文。
	CategoriesCn *string `json:"categories_cn,omitempty"`

	// **参数解释：** 检视意见模块。
	Modules *[]string `json:"modules,omitempty"`

	// **参数解释：** 严重程度key。
	ReviewSeverity *string `json:"review_severity,omitempty"`

	// **参数解释：** 是否设置为默认模板。
	IsDefault *bool `json:"is_default,omitempty"`

	// **参数解释：** 模板作者id。
	CreatorId *int32 `json:"creator_id,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	Creator *UserBasicDto `json:"creator,omitempty"`

	Assignee *UserBasicDto `json:"assignee,omitempty"`
}

func (o DiscussionTemplateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscussionTemplateDto struct{}"
	}

	return strings.Join([]string{"DiscussionTemplateDto", string(data)}, " ")
}
