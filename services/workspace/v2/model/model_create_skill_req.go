package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSkillReq 创建企业自研技能请求。
type CreateSkillReq struct {

	// 技能slug，创建后不可修改。
	Slug string `json:"slug"`

	// 技能名称。
	DisplayName string `json:"display_name"`

	// 别名（业务界面可修改的名称）。
	AliasName *string `json:"alias_name,omitempty"`

	// 技能描述。
	Description *string `json:"description,omitempty"`

	Category *SkillCategoryEnum `json:"category"`

	// 技能标签。
	Tags *[]string `json:"tags,omitempty"`

	// 封面图 base64 编码。
	Cover *string `json:"cover,omitempty"`

	Source *SkillSourceEnum `json:"source,omitempty"`

	// 支持的操作系统类型列表。
	SupportOsTypes *[]string `json:"support_os_types,omitempty"`

	// 技能包信息列表。
	Packages *[]CreateSkillPackage `json:"packages,omitempty"`
}

func (o CreateSkillReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillReq struct{}"
	}

	return strings.Join([]string{"CreateSkillReq", string(data)}, " ")
}
