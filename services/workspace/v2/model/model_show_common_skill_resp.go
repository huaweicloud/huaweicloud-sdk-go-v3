package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommonSkillResp 查询公共技能详情响应。
type ShowCommonSkillResp struct {

	// 技能id。
	Id *string `json:"id,omitempty"`

	// 技能slug。
	Slug *string `json:"slug,omitempty"`

	// 技能名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 别名。
	AliasName *string `json:"alias_name,omitempty"`

	// 技能描述。
	Description *string `json:"description,omitempty"`

	Category *SkillCategoryEnum `json:"category,omitempty"`

	// 技能标签。
	Tags *[]string `json:"tags,omitempty"`

	OwnerType *SkillOwnerTypeEnum `json:"owner_type,omitempty"`

	VisibilityScope *VisibilityScopeEnum `json:"visibility_scope,omitempty"`

	// 可见租户 ID 列表，仅 visibility_scope=SPECIFIC_TENANTS 时返回。
	VisibleDomainIds *[]string `json:"visible_domain_ids,omitempty"`

	Status *SkillStatusEnum `json:"status,omitempty"`

	// 当前生效的技能包id。
	CurrentPackageId *string `json:"current_package_id,omitempty"`

	// 技能封面图 base64 编码。
	Cover *string `json:"cover,omitempty"`

	Source *SkillSourceEnum `json:"source,omitempty"`

	// 支持的操作系统类型列表。
	SupportOsTypes *[]string `json:"support_os_types,omitempty"`

	// 技能包摘要列表。
	Packages *[]SkillPackageSummary `json:"packages,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`

	// 已绑定实例数量。
	AttachInstanceNumber *int32 `json:"attach_instance_number,omitempty"`
}

func (o ShowCommonSkillResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommonSkillResp struct{}"
	}

	return strings.Join([]string{"ShowCommonSkillResp", string(data)}, " ")
}
