package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSkillOrderListRequest Request Object
type ShowSkillOrderListRequest struct {

	// 每页显示的条目数量, 最大 100，默认值 10
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始位置, 默认值 0
	Offset *int32 `json:"offset,omitempty"`

	// 技能名称，支持模糊匹配。中英文、数字、下划线、中划线 长度[1-60]
	SkillName *string `json:"skill_name,omitempty"`

	// 技能形式，no_termplate不支持Modelbox部署模板，support_template支持Modelbox模板。
	SkillForm *string `json:"skill_form,omitempty"`
}

func (o ShowSkillOrderListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillOrderListRequest struct{}"
	}

	return strings.Join([]string{"ShowSkillOrderListRequest", string(data)}, " ")
}
