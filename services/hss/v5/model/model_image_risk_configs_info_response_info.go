package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageRiskConfigsInfoResponseInfo 配置检测结果信息
type ImageRiskConfigsInfoResponseInfo struct {

	// **参数解释** 镜像安全配置检测结果的风险等级，用于筛选指定风险等级的检测记录 **约束限制** 取值必须在指定范围内，否则返回空结果 **取值范围** - Security：安全 - Low：低危 - Medium：中危 - High：高危 **默认取值** 无
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 安全配置检测的基线名称，用于筛选指定基线的检测结果（如'CentOS 7'、'EulerOS'等） **约束限制** 仅支持功能介绍中列出的系统基线（CentOS 7、Debian 10、EulerOS、Ubuntu16） **取值范围** 支持的基线名称列表详见功能介绍 **默认取值** 无
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释** 用于区分基线的类型 **取值范围** 字符长度0-256位
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释** 安全配置检测遵循的标准，用于筛选符合指定标准的检测结果 **约束限制** 取值必须在指定范围内，否则返回空结果 **取值范围** - cn_standard：等保合规标准 - hw_standard：云安全实践标准 **默认取值** 无
	Standard *string `json:"standard,omitempty"`

	// **参数解释** 该基线对应的安全配置检测总检查项数量 **取值范围** 取值0-2097152
	CheckRuleNum *int32 `json:"check_rule_num,omitempty"`

	// **参数解释** 该基线检测中未通过（存在安全风险）的检查项数量 **取值范围** 取值0-2097152
	FailedRuleNum *int32 `json:"failed_rule_num,omitempty"`

	// **参数解释** 该基线的详细描述，说明基线的检测目的、适用场景等信息 **取值范围** 字符长度0-65534位，支持中文、英文、数字、常用标点符号及空格
	CheckTypeDesc *string `json:"check_type_desc,omitempty"`
}

func (o ImageRiskConfigsInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageRiskConfigsInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"ImageRiskConfigsInfoResponseInfo", string(data)}, " ")
}
