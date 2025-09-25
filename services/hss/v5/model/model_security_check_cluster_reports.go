package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SecurityCheckClusterReports 集群的安全体检结果信息
type SecurityCheckClusterReports struct {

	// **参数解释**： 体检报告ID **取值范围**： 不涉及
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群ID **取值范围**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群类型 **取值范围**： - CCE：CCE Standard集群 - Turbo：CCE Turbo集群
	ClusterCategory *SecurityCheckClusterReportsClusterCategory `json:"cluster_category,omitempty"`

	// **参数解释**： 风险级别 **取值范围**： - Security：无风险 - Insecurity：有风险
	Severity *SecurityCheckClusterReportsSeverity `json:"severity,omitempty"`

	// **参数解释**： 本地镜像漏洞风险数量 **取值范围**： 不涉及
	LocalImageVulNum *int32 `json:"local_image_vul_num,omitempty"`

	// **参数解释**： 基线风险数量 **取值范围**： 不涉及
	RiskConfigNum *int32 `json:"risk_config_num,omitempty"`

	// **参数解释**： 特权容器告警数量 **取值范围**： 不涉及
	PrivilegedContainerNum *int32 `json:"privileged_container_num,omitempty"`

	// **参数解释**： 最新检测时间 **取值范围**： 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`
}

func (o SecurityCheckClusterReports) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckClusterReports struct{}"
	}

	return strings.Join([]string{"SecurityCheckClusterReports", string(data)}, " ")
}

type SecurityCheckClusterReportsClusterCategory struct {
	value string
}

type SecurityCheckClusterReportsClusterCategoryEnum struct {
	CCE   SecurityCheckClusterReportsClusterCategory
	TURBO SecurityCheckClusterReportsClusterCategory
}

func GetSecurityCheckClusterReportsClusterCategoryEnum() SecurityCheckClusterReportsClusterCategoryEnum {
	return SecurityCheckClusterReportsClusterCategoryEnum{
		CCE: SecurityCheckClusterReportsClusterCategory{
			value: "CCE",
		},
		TURBO: SecurityCheckClusterReportsClusterCategory{
			value: "Turbo",
		},
	}
}

func (c SecurityCheckClusterReportsClusterCategory) Value() string {
	return c.value
}

func (c SecurityCheckClusterReportsClusterCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SecurityCheckClusterReportsClusterCategory) UnmarshalJSON(b []byte) error {
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

type SecurityCheckClusterReportsSeverity struct {
	value string
}

type SecurityCheckClusterReportsSeverityEnum struct {
	SECURITY   SecurityCheckClusterReportsSeverity
	INSECURITY SecurityCheckClusterReportsSeverity
}

func GetSecurityCheckClusterReportsSeverityEnum() SecurityCheckClusterReportsSeverityEnum {
	return SecurityCheckClusterReportsSeverityEnum{
		SECURITY: SecurityCheckClusterReportsSeverity{
			value: "Security",
		},
		INSECURITY: SecurityCheckClusterReportsSeverity{
			value: "Insecurity",
		},
	}
}

func (c SecurityCheckClusterReportsSeverity) Value() string {
	return c.value
}

func (c SecurityCheckClusterReportsSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SecurityCheckClusterReportsSeverity) UnmarshalJSON(b []byte) error {
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
