package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportBaselineRequestBody struct {

	// **参数解释**： 导出类型 **约束限制**： 不涉及 **取值范围**： - risk-config : 基线配置检测 - password-complexity : 口令复杂度检测 - weak-password : 弱口令检测  **默认取值**： risk-config
	Category *string `json:"category,omitempty"`

	// **参数解释** 导出数据条数 **约束限制** 不涉及 **取值范围**  取值 1 - 200000  **默认取值** 100000
	ExportSize *int32 `json:"export_size,omitempty"`

	// **参数解释**： 策略组id **约束限制**： 不涉及 **取值范围**： 字符长度0-128位  **默认取值**: 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 配置检查（基线）的名称，例如SSH、CentOS 7、Windows **约束限制**： 不涉及 **取值范围**： 字符长度0-256位  **默认取值**: 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释**： 标准类型 **约束限制**： 不涉及 **取值范围**： - cn_standard ： 等保合规标准 - hw_standard ： 云安全实践标准 - cis_standard ： 通用安全标准  **默认取值**: 不涉及
	Standard *string `json:"standard,omitempty"`

	// **参数解释**： 检查结果 **约束限制**： 不涉及 **取值范围** - pass : 检查通过 - failed : 检查未通过(已废弃) - unhandled: 未通过  **默认取值** 不涉及
	ScanResult *string `json:"scan_result,omitempty"`

	// **参数解释**： 风险等级 **约束限制**： 不涉及 **取值范围** - Security: 无风险 - Low : 低危 - Medium : 中危 - High: 高危  **默认取值** 不涉及
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 主机ID，不赋值时，查租户所有主机 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位  **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 导出配置检测数据的表头信息列表
	ExportHeaders [][]string `json:"export_headers"`

	// **参数解释** 基线检查项的类型 **约束限制** 不涉及 **取值范围** 字符长度0-256位 **默认取值** 不涉及
	Tag *string `json:"tag,omitempty"`

	// **参数解释** 配置检查（基线）的名称，例如SSH、CentOS 7、Windows **约束限制** 不涉及 **取值范围** 字符长度0-256位 **默认取值** 不涉及
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释** 统计结果类型，包含如下： **约束限制** 不涉及 **取值范围** - pass      : 已通过，表示查看主机全部通过的检查项 - failed    : 未通过，表示查看主机全部未通过 & 全部未处理的检查项 - processed : 已处理，表示查看主机存在未通过 & 未通过主机已全部处理(忽略、加白)的检查项  **默认取值** 不涉及
	StatisticsScanResult *string `json:"statistics_scan_result,omitempty"`

	// **参数解释** 检查项（检查规则）名称，支持模糊匹配 **约束限制** 不涉及 **取值范围** 字符长度0-2048位 **默认取值** 不涉及
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释** 集群ID **约束限制** 不涉及 **取值范围** 字符长度0-64位 **默认取值** 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ExportBaselineRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportBaselineRequestBody struct{}"
	}

	return strings.Join([]string{"ExportBaselineRequestBody", string(data)}, " ")
}
