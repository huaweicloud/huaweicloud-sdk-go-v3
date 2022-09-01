package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置检测结果信息
type SecurityCheckInfoResponseInfo struct {

	// 风险等级，包含如下:   - Low : 低危   - Medium : 中危   - High : 高危
	Severity *string `json:"severity,omitempty" xml:"severity"`

	// 基线名称
	CheckName *string `json:"check_name,omitempty" xml:"check_name"`

	// 基线类型
	CheckType *string `json:"check_type,omitempty" xml:"check_type"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 华为标准   - qt_standard : 青腾标准
	Standard *string `json:"standard,omitempty" xml:"standard"`

	// 检查项数量
	CheckRuleNum *int32 `json:"check_rule_num,omitempty" xml:"check_rule_num"`

	// 风险项数量
	FailedRuleNum *int32 `json:"failed_rule_num,omitempty" xml:"failed_rule_num"`

	// 影响的服务器数量
	HostNum *int32 `json:"host_num,omitempty" xml:"host_num"`

	// 最新检测时间
	ScanTime *int64 `json:"scan_time,omitempty" xml:"scan_time"`

	// 基线描述信息
	CheckTypeDesc *string `json:"check_type_desc,omitempty" xml:"check_type_desc"`
}

func (o SecurityCheckInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckInfoResponseInfo", string(data)}, " ")
}
