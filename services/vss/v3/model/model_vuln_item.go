package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type VulnItem struct {
	// 漏洞ID

	VulnId *string `json:"vuln_id,omitempty"`
	// 域名ID

	DomainId *string `json:"domain_id,omitempty"`
	// 目标网址

	Url *string `json:"url,omitempty"`
	// 漏洞风险等级:   * high - 高风险   * middle - 中风险   * low - 低风险   * hint - 提示

	Severity *VulnItemSeverity `json:"severity,omitempty"`
	// 漏洞状态:   * repairing - 未修复   * repaired - 已修复   * false_report - 误报，已忽略

	VulnStatus *VulnItemVulnStatus `json:"vuln_status,omitempty"`
	// 漏洞类别

	VulnClass *string `json:"vuln_class,omitempty"`
	// 漏洞名称

	VulnType *string `json:"vuln_type,omitempty"`
	// 漏洞描述

	Description *string `json:"description,omitempty"`
	// 修复建议

	Advice *string `json:"advice,omitempty"`
	// 命中详情

	HitDetails *string `json:"hit_details,omitempty"`
	// 请求详情

	Request *string `json:"request,omitempty"`
	// 响应详情

	Response *string `json:"response,omitempty"`
	// 漏洞确认人

	Provider *string `json:"provider,omitempty"`
	// 漏洞忽略理由

	Reason *string `json:"reason,omitempty"`
	// 漏洞发现时间

	FindTime *string `json:"find_time,omitempty"`
}

func (o VulnItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulnItem struct{}"
	}

	return strings.Join([]string{"VulnItem", string(data)}, " ")
}

type VulnItemSeverity struct {
	value string
}

type VulnItemSeverityEnum struct {
	HIGH   VulnItemSeverity
	MIDDLE VulnItemSeverity
	LOW    VulnItemSeverity
	HINT   VulnItemSeverity
}

func GetVulnItemSeverityEnum() VulnItemSeverityEnum {
	return VulnItemSeverityEnum{
		HIGH: VulnItemSeverity{
			value: "high",
		},
		MIDDLE: VulnItemSeverity{
			value: "middle",
		},
		LOW: VulnItemSeverity{
			value: "low",
		},
		HINT: VulnItemSeverity{
			value: "hint",
		},
	}
}

func (c VulnItemSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VulnItemSeverity) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type VulnItemVulnStatus struct {
	value string
}

type VulnItemVulnStatusEnum struct {
	REPAIRING    VulnItemVulnStatus
	REPAIRED     VulnItemVulnStatus
	FALSE_REPORT VulnItemVulnStatus
}

func GetVulnItemVulnStatusEnum() VulnItemVulnStatusEnum {
	return VulnItemVulnStatusEnum{
		REPAIRING: VulnItemVulnStatus{
			value: "repairing",
		},
		REPAIRED: VulnItemVulnStatus{
			value: "repaired",
		},
		FALSE_REPORT: VulnItemVulnStatus{
			value: "false_report",
		},
	}
}

func (c VulnItemVulnStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VulnItemVulnStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
