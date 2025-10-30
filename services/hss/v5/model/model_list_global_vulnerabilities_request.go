package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalVulnerabilitiesRequest Request Object
type ListGlobalVulnerabilitiesRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 修复紧急度 **取值范围**: - immediate_repair：需尽快修复。 - delay_repair：可延后修复。 - not_needed_repair：暂可不修复。  **约束限制**: 不涉及 **默认取值**: 不涉及
	RepairNecessity *ListGlobalVulnerabilitiesRequestRepairNecessity `json:"repair_necessity,omitempty"`

	// **参数解释**: 漏洞ID（支持模糊查询） **约束限制**: 不涉及 **取值范围**: 字符长度0-65535位 **默认取值**: 不涉及
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: - linux_vul：linux漏洞。 - app_vul：应用漏洞。  **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`
}

func (o ListGlobalVulnerabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalVulnerabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalVulnerabilitiesRequest", string(data)}, " ")
}

type ListGlobalVulnerabilitiesRequestRepairNecessity struct {
	value string
}

type ListGlobalVulnerabilitiesRequestRepairNecessityEnum struct {
	IMMEDIATE_REPAIR  ListGlobalVulnerabilitiesRequestRepairNecessity
	DELAY_REPAIR      ListGlobalVulnerabilitiesRequestRepairNecessity
	NOT_NEEDED_REPAIR ListGlobalVulnerabilitiesRequestRepairNecessity
}

func GetListGlobalVulnerabilitiesRequestRepairNecessityEnum() ListGlobalVulnerabilitiesRequestRepairNecessityEnum {
	return ListGlobalVulnerabilitiesRequestRepairNecessityEnum{
		IMMEDIATE_REPAIR: ListGlobalVulnerabilitiesRequestRepairNecessity{
			value: "immediate_repair",
		},
		DELAY_REPAIR: ListGlobalVulnerabilitiesRequestRepairNecessity{
			value: "delay_repair",
		},
		NOT_NEEDED_REPAIR: ListGlobalVulnerabilitiesRequestRepairNecessity{
			value: "not_needed_repair",
		},
	}
}

func (c ListGlobalVulnerabilitiesRequestRepairNecessity) Value() string {
	return c.value
}

func (c ListGlobalVulnerabilitiesRequestRepairNecessity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalVulnerabilitiesRequestRepairNecessity) UnmarshalJSON(b []byte) error {
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
