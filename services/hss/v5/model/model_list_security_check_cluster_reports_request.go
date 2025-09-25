package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityCheckClusterReportsRequest Request Object
type ListSecurityCheckClusterReportsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 集群名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群ID **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群类型 **约束限制**： 不涉及 **取值范围**： - CCE：CCE Standard集群 - Turbo：CCE Turbo集群  **默认取值**： 不涉及
	ClusterCategory *ListSecurityCheckClusterReportsRequestClusterCategory `json:"cluster_category,omitempty"`

	// **参数解释**： 风险级别 **约束限制**： 不涉及 **取值范围**： - Security：无风险 - Insecurity：有风险  **默认取值**： 不涉及
	Severity *ListSecurityCheckClusterReportsRequestSeverity `json:"severity,omitempty"`
}

func (o ListSecurityCheckClusterReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityCheckClusterReportsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityCheckClusterReportsRequest", string(data)}, " ")
}

type ListSecurityCheckClusterReportsRequestClusterCategory struct {
	value string
}

type ListSecurityCheckClusterReportsRequestClusterCategoryEnum struct {
	CCE   ListSecurityCheckClusterReportsRequestClusterCategory
	TURBO ListSecurityCheckClusterReportsRequestClusterCategory
}

func GetListSecurityCheckClusterReportsRequestClusterCategoryEnum() ListSecurityCheckClusterReportsRequestClusterCategoryEnum {
	return ListSecurityCheckClusterReportsRequestClusterCategoryEnum{
		CCE: ListSecurityCheckClusterReportsRequestClusterCategory{
			value: "CCE",
		},
		TURBO: ListSecurityCheckClusterReportsRequestClusterCategory{
			value: "Turbo",
		},
	}
}

func (c ListSecurityCheckClusterReportsRequestClusterCategory) Value() string {
	return c.value
}

func (c ListSecurityCheckClusterReportsRequestClusterCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityCheckClusterReportsRequestClusterCategory) UnmarshalJSON(b []byte) error {
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

type ListSecurityCheckClusterReportsRequestSeverity struct {
	value string
}

type ListSecurityCheckClusterReportsRequestSeverityEnum struct {
	SECURITY   ListSecurityCheckClusterReportsRequestSeverity
	INSECURITY ListSecurityCheckClusterReportsRequestSeverity
}

func GetListSecurityCheckClusterReportsRequestSeverityEnum() ListSecurityCheckClusterReportsRequestSeverityEnum {
	return ListSecurityCheckClusterReportsRequestSeverityEnum{
		SECURITY: ListSecurityCheckClusterReportsRequestSeverity{
			value: "Security",
		},
		INSECURITY: ListSecurityCheckClusterReportsRequestSeverity{
			value: "Insecurity",
		},
	}
}

func (c ListSecurityCheckClusterReportsRequestSeverity) Value() string {
	return c.value
}

func (c ListSecurityCheckClusterReportsRequestSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityCheckClusterReportsRequestSeverity) UnmarshalJSON(b []byte) error {
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
