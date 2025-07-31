package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineDirectoryRequest Request Object
type ShowBaselineDirectoryRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 基线检查的操作系统 **约束限制** 不涉及 **取值范围** - Linux - Windows  **默认取值** Linux
	SupportOs string `json:"support_os"`

	// **参数解释** 决定目录结构的顺序 **约束限制** 不涉及 **取值范围** - check_type : 二级目录为基线名称 - tag        : 二级目录为检查项的类型  **默认取值** 不涉及
	SelectType string `json:"select_type"`

	// **参数解释** 展示该策略组选中哪些检查项 **约束限制** 不涉及 **取值范围** 字符长度0-64位 **默认取值** 不涉及
	GroupId *string `json:"group_id,omitempty"`
}

func (o ShowBaselineDirectoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineDirectoryRequest struct{}"
	}

	return strings.Join([]string{"ShowBaselineDirectoryRequest", string(data)}, " ")
}
