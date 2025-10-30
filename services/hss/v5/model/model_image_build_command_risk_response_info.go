package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageBuildCommandRiskResponseInfo 镜像构建指令风险信息
type ImageBuildCommandRiskResponseInfo struct {

	// **参数解释**: 最后一次检测时间，时间单位 毫秒（ms） **取值范围**: 大小0-9223372036854775807
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 受影响的镜像的数量，进行了当前构建指令的镜像数量 **取值范围**: 大小0-2097152
	ImageNum *int32 `json:"image_num,omitempty"`

	// **参数解释** 风险id **取值范围** 字符长度1-64位
	RiskId *string `json:"risk_id,omitempty"`

	// **参数解释** 风险检测规则id **取值范围** 字符长度1-64位
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释** 风险名称 **取值范围** 字符长度1-512位
	RiskName *string `json:"risk_name,omitempty"`

	// **参数解释** 风险程度 **取值范围** - critical：严重。 - high：高危。 - medium：中危。 - low：低危。
	RiskLevel *string `json:"risk_level,omitempty"`

	// **参数解释** 风险描述 **取值范围** 字符长度0-65534位
	Description *string `json:"description,omitempty"`

	// **参数解释** 镜像类型 **取值范围** 字符长度1-32位
	ImageType *string `json:"image_type,omitempty"`
}

func (o ImageBuildCommandRiskResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageBuildCommandRiskResponseInfo struct{}"
	}

	return strings.Join([]string{"ImageBuildCommandRiskResponseInfo", string(data)}, " ")
}
