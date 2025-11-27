package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRuleRiskInfoResponseInfo 检查项风险信息
type CheckRuleRiskInfoResponseInfo struct {

	// **参数解释**: 风险等级 **取值范围**: - Low：低危 - Medium：中危 - High：高危
	Severity *string `json:"severity,omitempty"`

	// **参数解释**: 配置检查（基线）的名称，例如SSH、CentOS 7、Windows **取值范围**: 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释**: 配置检查（基线）的类型,Linux系统支持的基线一般check_type和check_name相同,例如SSH、CentOS 7。 Windows系统支持的基线一般check_type和check_name不相同，例如check_name为Windows的配置检查（基线），它的check_type包含Windows Server 2019 R2、Windows Server 2016 R2等。 **取值范围**: 不涉及
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释**: 标准类型 **取值范围**: - cn_standard：等保合规标准 - hw_standard：云安全实践标准
	Standard *string `json:"standard,omitempty"`

	// **参数解释**: 检查项（检查规则）名称 **取值范围**: 不涉及
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释**: 检查项ID **取值范围**: 不涉及
	CheckRuleId *string `json:"check_rule_id,omitempty"`

	// **参数解释**: 受影响的服务器的数量，进行了当前基线检测的服务器数量 **取值范围**: 不涉及
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 检测结果 **取值范围**: - pass : 检测通过 - failed : 检测不通过
	ScanResult *string `json:"scan_result,omitempty"`

	// **参数解释**: 检测项状态 **取值范围**: - safe : 无需处理 - ignored : 已忽略 - unhandled : 未处理 - fixing : 修复中 - fix-failed : 修复失败 - verifying : 验证中
	Status *string `json:"status,omitempty"`

	// **参数解释**: 是否支持一键修复 **取值范围**: - 1：支持一键修复 - 0：不支持
	EnableFix *int32 `json:"enable_fix,omitempty"`

	// **参数解释**: 该检查项的修复&忽略&验证按钮是否可单击 **取值范围**: - true：按钮可单击 - false：按钮不可单击
	EnableClick *bool `json:"enable_click,omitempty"`

	// **参数解释** 不可点击的原因 **取值范围**  字符长度0-512位
	NotEnableClickDescription *string `json:"not_enable_click_description,omitempty"`

	// **参数解释**: 支持传递参数修复的检查项可传递参数的范围，只有支持传递参数修复的检查项才返回此数据 **取值范围**: 不涉及
	RuleParams *[]CheckRuleFixParamInfo `json:"rule_params,omitempty"`
}

func (o CheckRuleRiskInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleRiskInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleRiskInfoResponseInfo", string(data)}, " ")
}
