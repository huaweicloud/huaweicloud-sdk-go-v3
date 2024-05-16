package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HostVulnItem struct {

	// 漏洞类型
	Type *string `json:"type,omitempty"`

	// 扫描端口号
	Port *string `json:"port,omitempty"`

	// 漏洞标题
	Title *string `json:"title,omitempty"`

	// 漏洞公告ID
	SaId *string `json:"sa_id,omitempty"`

	// 漏洞ID
	VulnId *string `json:"vuln_id,omitempty"`

	// 漏洞风险等级:   * high - 高风险   * medium - 中风险   * low - 低风险   * hint - 提示
	Severity *HostVulnItemSeverity `json:"severity,omitempty"`

	// 漏洞摘要
	Topic *string `json:"topic,omitempty"`

	// 漏洞描述
	Description *string `json:"description,omitempty"`

	// 漏洞提示建议
	Solution *string `json:"solution,omitempty"`

	// 漏洞修复建议
	FixAdvisory *string `json:"fix_advisory,omitempty"`

	// 漏洞修复指令
	FixCmd *string `json:"fix_cmd,omitempty"`

	// CVE漏洞列表
	CveList *[]HostVulnItemCveList `json:"cve_list,omitempty"`

	// 参考信息链接列表
	RefLinkList *[]string `json:"ref_link_list,omitempty"`

	// 内容列表
	ComponentList *[]HostVulnItemComponentList `json:"component_list,omitempty"`

	// 检查结果
	VulDetectResult *string `json:"vul_detect_result,omitempty"`

	// CVSS分数信息
	CvssScore *string `json:"cvss_score,omitempty"`

	// CVSS版本信息
	CvssVersion *string `json:"cvss_version,omitempty"`

	// CVSS向量信息
	CvssVector *string `json:"cvss_vector,omitempty"`

	// 是否误报
	IsIgnore *bool `json:"is_ignore,omitempty"`
}

func (o HostVulnItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulnItem struct{}"
	}

	return strings.Join([]string{"HostVulnItem", string(data)}, " ")
}

type HostVulnItemSeverity struct {
	value string
}

type HostVulnItemSeverityEnum struct {
	HIGH   HostVulnItemSeverity
	MEDIUM HostVulnItemSeverity
	LOW    HostVulnItemSeverity
	HINT   HostVulnItemSeverity
}

func GetHostVulnItemSeverityEnum() HostVulnItemSeverityEnum {
	return HostVulnItemSeverityEnum{
		HIGH: HostVulnItemSeverity{
			value: "high",
		},
		MEDIUM: HostVulnItemSeverity{
			value: "medium",
		},
		LOW: HostVulnItemSeverity{
			value: "low",
		},
		HINT: HostVulnItemSeverity{
			value: "hint",
		},
	}
}

func (c HostVulnItemSeverity) Value() string {
	return c.value
}

func (c HostVulnItemSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostVulnItemSeverity) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
