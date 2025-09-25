package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHandleAffectBaselineRequestBody **参数解释** 基线检查执行操作时影响的范围的请求参数
type ListHandleAffectBaselineRequestBody struct {

	// **参数解释** 基线检查执行的操作 **约束限制** 不涉及 **取值范围** - add_to_whitelist: 加白名单 - ignore          : 忽略 - unignore        : 取消忽略 - fix             : 修复 - verify          : 验证 **默认取值** 不涉及
	Action string `json:"action"`

	// **参数解释** 当前检查项的状态 **约束限制** 不涉及 **取值范围** - unhandled : 未处理 - fix-failed: 修复失败 - fixing    : 修复中 - verifying : 验证中 - ignored   : 忽略 - safe      : 安全 **默认取值** 不涉及
	HandleStatus string `json:"handle_status"`

	// **参数解释** 主机id，没有该字段则代表该检查项影响的部分主机 **约束限制** 不涉及 **取值范围** 字符长度1-256位 **默认取值** 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 需要进行操作的检查项列表 **约束限制** 列表范围0-200条
	CheckRuleList []ListHandleAffectBaselineRequestBodyCheckRuleList `json:"check_rule_list"`
}

func (o ListHandleAffectBaselineRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHandleAffectBaselineRequestBody struct{}"
	}

	return strings.Join([]string{"ListHandleAffectBaselineRequestBody", string(data)}, " ")
}
