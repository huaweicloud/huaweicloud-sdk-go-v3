package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClusterSecurityCheckBaseLineDetectionInfo 配置检测信息
type ClusterSecurityCheckBaseLineDetectionInfo struct {

	// **参数解释**： 基线风险级别 **取值范围**： - High：高危基线 - Medium：中危基线 - Low：低危基线
	Severity *ClusterSecurityCheckBaseLineDetectionInfoSeverity `json:"severity,omitempty"`

	// **参数解释**： 基线名称 **取值范围**： 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释**： 基线类型 **取值范围**： 不涉及
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释**： 标准类型 **取值范围**： - hw_standard：云安全实践 - cn_standard：等保合规 - cis_standard：通用安全标准
	Standard *string `json:"standard,omitempty"`

	// **参数解释**： 检查项数量 **取值范围**： 不涉及
	CheckRuleNum *int32 `json:"check_rule_num,omitempty"`

	// **参数解释**： 风险项数量 **取值范围**： 不涉及
	FailedRuleNum *int32 `json:"failed_rule_num,omitempty"`

	// **参数解释**： 基线描述信息 **取值范围**： 不涉及
	CheckTypeDesc *string `json:"check_type_desc,omitempty"`

	// **参数解释**： 基线检测列表 **取值范围**： 不涉及
	BaselineItemList *[]ClusterSecurityCheckBaselineItemInfo `json:"baseline_item_list,omitempty"`
}

func (o ClusterSecurityCheckBaseLineDetectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSecurityCheckBaseLineDetectionInfo struct{}"
	}

	return strings.Join([]string{"ClusterSecurityCheckBaseLineDetectionInfo", string(data)}, " ")
}

type ClusterSecurityCheckBaseLineDetectionInfoSeverity struct {
	value string
}

type ClusterSecurityCheckBaseLineDetectionInfoSeverityEnum struct {
	HIGH   ClusterSecurityCheckBaseLineDetectionInfoSeverity
	MEDIUM ClusterSecurityCheckBaseLineDetectionInfoSeverity
	LOW    ClusterSecurityCheckBaseLineDetectionInfoSeverity
}

func GetClusterSecurityCheckBaseLineDetectionInfoSeverityEnum() ClusterSecurityCheckBaseLineDetectionInfoSeverityEnum {
	return ClusterSecurityCheckBaseLineDetectionInfoSeverityEnum{
		HIGH: ClusterSecurityCheckBaseLineDetectionInfoSeverity{
			value: "High",
		},
		MEDIUM: ClusterSecurityCheckBaseLineDetectionInfoSeverity{
			value: "Medium",
		},
		LOW: ClusterSecurityCheckBaseLineDetectionInfoSeverity{
			value: "Low",
		},
	}
}

func (c ClusterSecurityCheckBaseLineDetectionInfoSeverity) Value() string {
	return c.value
}

func (c ClusterSecurityCheckBaseLineDetectionInfoSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterSecurityCheckBaseLineDetectionInfoSeverity) UnmarshalJSON(b []byte) error {
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
