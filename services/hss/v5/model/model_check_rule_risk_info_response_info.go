package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 检查项风险信息
type CheckRuleRiskInfoResponseInfo struct {

	// 风险等级，包含如下:   - Low : 低危   - Medium : 中危   - High : 高危
	Severity *string `json:"severity,omitempty" xml:"severity"`

	// 基线名称
	CheckType *string `json:"check_type,omitempty" xml:"check_type"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 华为标准   - qt_standard : 青腾标准
	Standard *string `json:"standard,omitempty" xml:"standard"`

	// 检查项
	CheckRuleName *string `json:"check_rule_name,omitempty" xml:"check_rule_name"`

	// 检查项ID
	CheckRuleId *string `json:"check_rule_id,omitempty" xml:"check_rule_id"`

	// 影响服务器数量
	HostNum *int32 `json:"host_num,omitempty" xml:"host_num"`

	// 检测结果，包含如下：   - pass   - failed
	ScanResult *string `json:"scan_result,omitempty" xml:"scan_result"`

	// 状态，包含如下：   - safe : 无需处理   - ignored : 已忽略   - unhandled : 未处理
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o CheckRuleRiskInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleRiskInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleRiskInfoResponseInfo", string(data)}, " ")
}
