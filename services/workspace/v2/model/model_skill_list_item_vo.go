package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SkillListItemVo 技能列表项响应。
type SkillListItemVo struct {

	// 技能id。
	Id *string `json:"id,omitempty"`

	// 技能slug。
	Slug *string `json:"slug,omitempty"`

	// 技能名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 别名。
	AliasName *string `json:"alias_name,omitempty"`

	Category *SkillCategoryEnum `json:"category,omitempty"`

	// 技能描述。
	Description *string `json:"description,omitempty"`

	// 技能标签。
	Tags *[]string `json:"tags,omitempty"`

	Status *SkillStatusEnum `json:"status,omitempty"`

	VisibilityScope *VisibilityScopeEnum `json:"visibility_scope,omitempty"`

	// 当前生效的技能包id。
	CurrentPackageId *string `json:"current_package_id,omitempty"`

	// 当前生效版本号。
	CurrentVersion *string `json:"current_version,omitempty"`

	// 当前生效版本修订号。
	CurrentRevision *int32 `json:"current_revision,omitempty"`

	// 技能封面图 base64 编码。
	Cover *string `json:"cover,omitempty"`

	Source *SkillSourceEnum `json:"source,omitempty"`

	// 支持的操作系统类型列表。
	SupportOsTypes *[]string `json:"support_os_types,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`

	// 已绑定实例数量。
	AttachInstanceNumber *int32 `json:"attach_instance_number,omitempty"`
}

func (o SkillListItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkillListItemVo struct{}"
	}

	return strings.Join([]string{"SkillListItemVo", string(data)}, " ")
}
