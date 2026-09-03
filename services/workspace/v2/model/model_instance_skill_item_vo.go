package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceSkillItemVo 实例绑定的技能列表项响应。
type InstanceSkillItemVo struct {

	// 技能id。
	SkillId *string `json:"skill_id,omitempty"`

	// 技能名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 技能slug。
	Slug *string `json:"slug,omitempty"`

	// 别名。
	AliasName *string `json:"alias_name,omitempty"`

	// 技能描述。
	Description *string `json:"description,omitempty"`

	OwnerType *SkillOwnerTypeEnum `json:"owner_type,omitempty"`

	InstallStatus *InstallStatusEnum `json:"install_status,omitempty"`

	// 安装的技能包id。
	PackageId *string `json:"package_id,omitempty"`

	// 安装的技能包版本号。
	Version *string `json:"version,omitempty"`

	// 安装时间（ISO8601格式，UTC时区）。
	InstalledAt *string `json:"installed_at,omitempty"`

	// 技能封面图 base64 编码。
	Cover *string `json:"cover,omitempty"`
}

func (o InstanceSkillItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSkillItemVo struct{}"
	}

	return strings.Join([]string{"InstanceSkillItemVo", string(data)}, " ")
}
