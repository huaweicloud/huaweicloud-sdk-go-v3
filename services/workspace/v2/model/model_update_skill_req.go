package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSkillReq 更新企业自研技能请求。
type UpdateSkillReq struct {

	// 技能名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 别名（业务界面可修改的名称）。
	AliasName *string `json:"alias_name,omitempty"`

	// 技能描述。
	Description *string `json:"description,omitempty"`

	Category *SkillCategoryEnum `json:"category,omitempty"`

	// 技能标签。
	Tags *[]string `json:"tags,omitempty"`

	// 封面图 base64 编码。
	Cover *string `json:"cover,omitempty"`

	Status *SkillStatusEnum `json:"status,omitempty"`

	Source *SkillSourceEnum `json:"source,omitempty"`

	// 支持的操作系统类型列表。
	SupportOsTypes *[]string `json:"support_os_types,omitempty"`
}

func (o UpdateSkillReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSkillReq struct{}"
	}

	return strings.Join([]string{"UpdateSkillReq", string(data)}, " ")
}
