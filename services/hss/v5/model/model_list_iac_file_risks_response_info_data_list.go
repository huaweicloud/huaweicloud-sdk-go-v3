package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIacFileRisksResponseInfoDataList 文件风险信息
type ListIacFileRisksResponseInfoDataList struct {

	// 风险id
	RiskId *string `json:"risk_id,omitempty"`

	// 风险检测规则id
	RuleId *string `json:"rule_id,omitempty"`

	// 风险名称
	RiskName *string `json:"risk_name,omitempty"`

	// 风险程度，包含如下   - high ：高危   - medium ：中危   - low ：低危   - tips ：提示
	RiskLevel *string `json:"risk_level,omitempty"`

	// 风险分类，包含如下：   - control_plane：控制平面   - access_control：访问控制   - network：网络   - workload：工作负载   - secrets：密钥管理   - node_escape：节点逃逸
	RiskCategory *string `json:"risk_category,omitempty"`

	// 风险数量
	RiskNum *int32 `json:"risk_num,omitempty"`

	// 最近一次扫描时间
	LastScanTime *int64 `json:"last_scan_time,omitempty"`

	// 风险描述
	Description *string `json:"description,omitempty"`

	// 风险的处置建议
	Remediation *string `json:"remediation,omitempty"`

	// 存在风险的构建指令
	BuildInstruction *string `json:"build_instruction,omitempty"`
}

func (o ListIacFileRisksResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIacFileRisksResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"ListIacFileRisksResponseInfoDataList", string(data)}, " ")
}
