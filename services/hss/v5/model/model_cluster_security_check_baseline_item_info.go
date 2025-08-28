package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClusterSecurityCheckBaselineItemInfo 配置检查项信息
type ClusterSecurityCheckBaselineItemInfo struct {

	// **参数解释**： 检查项风险等级 **取值范围**： - High：高危 - Medium：中危 - Low：低危
	Severity *ClusterSecurityCheckBaselineItemInfoSeverity `json:"severity,omitempty"`

	// **参数解释**： 检查项 **取值范围**： 不涉及
	CheckItem *string `json:"check_item,omitempty"`

	// **参数解释**： 检查描述 **取值范围**： 不涉及
	CheckDescription *string `json:"check_description,omitempty"`

	// **参数解释**： 修复建议 **取值范围**： 不涉及
	CheckRemediation *string `json:"check_remediation,omitempty"`
}

func (o ClusterSecurityCheckBaselineItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSecurityCheckBaselineItemInfo struct{}"
	}

	return strings.Join([]string{"ClusterSecurityCheckBaselineItemInfo", string(data)}, " ")
}

type ClusterSecurityCheckBaselineItemInfoSeverity struct {
	value string
}

type ClusterSecurityCheckBaselineItemInfoSeverityEnum struct {
	HIGH   ClusterSecurityCheckBaselineItemInfoSeverity
	MEDIUM ClusterSecurityCheckBaselineItemInfoSeverity
	LOW    ClusterSecurityCheckBaselineItemInfoSeverity
}

func GetClusterSecurityCheckBaselineItemInfoSeverityEnum() ClusterSecurityCheckBaselineItemInfoSeverityEnum {
	return ClusterSecurityCheckBaselineItemInfoSeverityEnum{
		HIGH: ClusterSecurityCheckBaselineItemInfoSeverity{
			value: "High",
		},
		MEDIUM: ClusterSecurityCheckBaselineItemInfoSeverity{
			value: "Medium",
		},
		LOW: ClusterSecurityCheckBaselineItemInfoSeverity{
			value: "Low",
		},
	}
}

func (c ClusterSecurityCheckBaselineItemInfoSeverity) Value() string {
	return c.value
}

func (c ClusterSecurityCheckBaselineItemInfoSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterSecurityCheckBaselineItemInfoSeverity) UnmarshalJSON(b []byte) error {
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
