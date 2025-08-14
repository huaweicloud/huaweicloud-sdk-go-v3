package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBaselineWhiteListsRequest Request Object
type ListBaselineWhiteListsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// 基线检查的检查项名称
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// 基线检查的操作系统 - Linux - Windows
	OsType *string `json:"os_type,omitempty"`

	// 基线检查白名单的规则范围 - specific_host：部分主机 - all_host：全部主机
	RuleType *string `json:"rule_type,omitempty"`

	// 基线检查中检查项的检查类型 - 访问控制 - 服务配置
	Tag *string `json:"tag,omitempty"`

	// 基线检查白名单备注
	Description *string `json:"description,omitempty"`
}

func (o ListBaselineWhiteListsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaselineWhiteListsRequest struct{}"
	}

	return strings.Join([]string{"ListBaselineWhiteListsRequest", string(data)}, " ")
}
