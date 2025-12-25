package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ManualVulScanRequestInfo 手工检测漏洞接口请求体
type ManualVulScanRequestInfo struct {

	// **参数解释**: 扫描的漏洞类型列表 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值200 **默认取值**: 不涉及
	ManualScanType *[]ManualVulScanRequestInfoManualScanType `json:"manual_scan_type,omitempty"`

	// **参数解释**: 是否为批量扫描操作 **约束限制**: 不涉及 **取值范围**:   - true：是批量扫描操作   - false：不是批量扫描操作  **默认取值**: 不涉及
	BatchFlag *bool `json:"batch_flag,omitempty"`

	// **参数解释**: 扫描主机的范围 **约束限制**: 不涉及 **取值范围**:   - all_host：扫描全部主机，此类型不需要填写agent_id_list   - specific_host：扫描指定主机  **默认取值**: 不涉及
	RangeType *string `json:"range_type,omitempty"`

	// **参数解释**: 需要扫描主机的agent id列表 **约束限制**: range_type值为specific_host时有效 **取值范围**: 最小值1，最大值200 **默认取值**: 不涉及
	AgentIdList *[]string `json:"agent_id_list,omitempty"`

	// **参数解释**: 扫描的应急漏洞id列表，该字段不传值则扫描所有应急漏洞 **约束限制**: manual_scan_type字段的值中包含“urgent_vul”时有效 **取值范围**: 最小值1，最大值200 **默认取值**: 不涉及
	UrgentVulIdList *[]string `json:"urgent_vul_id_list,omitempty"`
}

func (o ManualVulScanRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManualVulScanRequestInfo struct{}"
	}

	return strings.Join([]string{"ManualVulScanRequestInfo", string(data)}, " ")
}

type ManualVulScanRequestInfoManualScanType struct {
	value string
}

type ManualVulScanRequestInfoManualScanTypeEnum struct {
	LINUX_VUL   ManualVulScanRequestInfoManualScanType
	WINDOWS_VUL ManualVulScanRequestInfoManualScanType
	WEB_CMS     ManualVulScanRequestInfoManualScanType
	APP_VUL     ManualVulScanRequestInfoManualScanType
	URGENT_VUL  ManualVulScanRequestInfoManualScanType
}

func GetManualVulScanRequestInfoManualScanTypeEnum() ManualVulScanRequestInfoManualScanTypeEnum {
	return ManualVulScanRequestInfoManualScanTypeEnum{
		LINUX_VUL: ManualVulScanRequestInfoManualScanType{
			value: "linux_vul",
		},
		WINDOWS_VUL: ManualVulScanRequestInfoManualScanType{
			value: "windows_vul",
		},
		WEB_CMS: ManualVulScanRequestInfoManualScanType{
			value: "web_cms",
		},
		APP_VUL: ManualVulScanRequestInfoManualScanType{
			value: "app_vul",
		},
		URGENT_VUL: ManualVulScanRequestInfoManualScanType{
			value: "urgent_vul",
		},
	}
}

func (c ManualVulScanRequestInfoManualScanType) Value() string {
	return c.value
}

func (c ManualVulScanRequestInfoManualScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ManualVulScanRequestInfoManualScanType) UnmarshalJSON(b []byte) error {
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
