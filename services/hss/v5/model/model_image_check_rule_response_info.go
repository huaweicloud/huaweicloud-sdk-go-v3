package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageCheckRuleResponseInfo 检查项风险信息
type ImageCheckRuleResponseInfo struct {

	// **参数解释**： 风险等级 **取值范围**： - Security：安全 - Low：低危 - Medium：中危 - High：高危
	Severity *string `json:"severity,omitempty"`

	// **参数解释**: 基线名称 **取值范围**: 字符长度0-256
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释**: 基线类型 **取值范围**: 字符长度0-256
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释**: 标准类型 **取值范围**: - cn_standard：等保合规标准 - hw_standard：华为标准 - qt_standard：青腾标准
	Standard *string `json:"standard,omitempty"`

	// **参数解释**: 基线描述信息 **取值范围**: 字符长度0-65534
	CheckTypeDesc *string `json:"check_type_desc,omitempty"`

	// **参数解释**: 检查项 **取值范围**: 字符长度0-2048
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// **参数解释**: 检查项ID **取值范围**: 字符长度0-128
	CheckRuleId *string `json:"check_rule_id,omitempty"`

	// **参数解释**: 检测结果 **取值范围**: - pass：通过 - failed：未通过
	ScanResult *string `json:"scan_result,omitempty"`

	// **参数解释**: 最后一次检测时间，时间单位：毫秒（ms） **取值范围**: 大小0-9223372036854775807
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 受影响的镜像的数量，进行了当前基线检测的镜像数量 **取值范围**: 大小0-2097152
	ImageNum *int32 `json:"image_num,omitempty"`
}

func (o ImageCheckRuleResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageCheckRuleResponseInfo struct{}"
	}

	return strings.Join([]string{"ImageCheckRuleResponseInfo", string(data)}, " ")
}
