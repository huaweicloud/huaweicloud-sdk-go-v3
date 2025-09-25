package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportHandledVulnerabilitiesRequestBody struct {

	// 漏洞名称
	VulName *string `json:"vul_name,omitempty"`

	// 漏洞修复优先级,包含如下 - Critical 紧急 - High 高 - Medium 中 - Low 低
	RepairPriority *ExportHandledVulnerabilitiesRequestBodyRepairPriority `json:"repair_priority,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 主机ip
	HostIp *string `json:"host_ip,omitempty"`

	// 是否影响业务
	IsAffectBusiness *bool `json:"is_affect_business,omitempty"`

	// 漏洞状态   - vul_status_unfix : 未处理   - vul_status_ignored : 已忽略   - vul_status_verified : 验证中   - vul_status_fixing : 修复中   - vul_status_fixed : 修复成功   - vul_status_reboot : 修复成功待重启   - vul_status_failed : 修复失败   - vul_status_fix_after_reboot : 请重启主机再次修复
	Status *ExportHandledVulnerabilitiesRequestBodyStatus `json:"status,omitempty"`

	// 资产重要性 important:重要 common：一般 test：测试
	AssetValue *ExportHandledVulnerabilitiesRequestBodyAssetValue `json:"asset_value,omitempty"`

	// 漏洞标签
	Label *string `json:"label,omitempty"`

	// 漏洞类型，包含如下：   -linux_vul : linux漏洞   -windows_vul : windows漏洞   -web_cms : Web-CMS漏洞   -app_vul : 应用漏洞   -urgent_vul : 应急漏洞
	Type *ExportHandledVulnerabilitiesRequestBodyType `json:"type,omitempty"`

	// 服务器组名称
	GroupName *string `json:"group_name,omitempty"`

	// 需要查询的漏洞处置周期：   - today：查询今日处理的数据   - all：查询所有已处理数据
	HandleCycle *ExportHandledVulnerabilitiesRequestBodyHandleCycle `json:"handle_cycle,omitempty"`

	// 指定要导出的漏洞数据
	SpecificVuls *[]ExportHandledVulnerabilitiesRequestBodySpecificVuls `json:"specific_vuls,omitempty"`

	// 导出数据条数
	ExportSize *int32 `json:"export_size,omitempty"`

	// 导出漏洞数据的表头信息列表
	ExportHeaders [][]string `json:"export_headers"`
}

func (o ExportHandledVulnerabilitiesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportHandledVulnerabilitiesRequestBody struct{}"
	}

	return strings.Join([]string{"ExportHandledVulnerabilitiesRequestBody", string(data)}, " ")
}

type ExportHandledVulnerabilitiesRequestBodyRepairPriority struct {
	value string
}

type ExportHandledVulnerabilitiesRequestBodyRepairPriorityEnum struct {
	CRITICAL ExportHandledVulnerabilitiesRequestBodyRepairPriority
	HIGH     ExportHandledVulnerabilitiesRequestBodyRepairPriority
	MEDIUM   ExportHandledVulnerabilitiesRequestBodyRepairPriority
	LOW      ExportHandledVulnerabilitiesRequestBodyRepairPriority
}

func GetExportHandledVulnerabilitiesRequestBodyRepairPriorityEnum() ExportHandledVulnerabilitiesRequestBodyRepairPriorityEnum {
	return ExportHandledVulnerabilitiesRequestBodyRepairPriorityEnum{
		CRITICAL: ExportHandledVulnerabilitiesRequestBodyRepairPriority{
			value: "Critical",
		},
		HIGH: ExportHandledVulnerabilitiesRequestBodyRepairPriority{
			value: "High",
		},
		MEDIUM: ExportHandledVulnerabilitiesRequestBodyRepairPriority{
			value: "Medium",
		},
		LOW: ExportHandledVulnerabilitiesRequestBodyRepairPriority{
			value: "Low",
		},
	}
}

func (c ExportHandledVulnerabilitiesRequestBodyRepairPriority) Value() string {
	return c.value
}

func (c ExportHandledVulnerabilitiesRequestBodyRepairPriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHandledVulnerabilitiesRequestBodyRepairPriority) UnmarshalJSON(b []byte) error {
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

type ExportHandledVulnerabilitiesRequestBodyStatus struct {
	value string
}

type ExportHandledVulnerabilitiesRequestBodyStatusEnum struct {
	VUL_STATUS_UNFIX            ExportHandledVulnerabilitiesRequestBodyStatus
	VUL_STATUS_IGNORED          ExportHandledVulnerabilitiesRequestBodyStatus
	VUL_STATUS_VERIFIED         ExportHandledVulnerabilitiesRequestBodyStatus
	VUL_STATUS_FIXING           ExportHandledVulnerabilitiesRequestBodyStatus
	VUL_STATUS_FIXED            ExportHandledVulnerabilitiesRequestBodyStatus
	VUL_STATUS_REBOOT           ExportHandledVulnerabilitiesRequestBodyStatus
	VUL_STATUS_FAILED           ExportHandledVulnerabilitiesRequestBodyStatus
	VUL_STATUS_FIX_AFTER_REBOOT ExportHandledVulnerabilitiesRequestBodyStatus
}

func GetExportHandledVulnerabilitiesRequestBodyStatusEnum() ExportHandledVulnerabilitiesRequestBodyStatusEnum {
	return ExportHandledVulnerabilitiesRequestBodyStatusEnum{
		VUL_STATUS_UNFIX: ExportHandledVulnerabilitiesRequestBodyStatus{
			value: "vul_status_unfix",
		},
		VUL_STATUS_IGNORED: ExportHandledVulnerabilitiesRequestBodyStatus{
			value: "vul_status_ignored",
		},
		VUL_STATUS_VERIFIED: ExportHandledVulnerabilitiesRequestBodyStatus{
			value: "vul_status_verified",
		},
		VUL_STATUS_FIXING: ExportHandledVulnerabilitiesRequestBodyStatus{
			value: "vul_status_fixing",
		},
		VUL_STATUS_FIXED: ExportHandledVulnerabilitiesRequestBodyStatus{
			value: "vul_status_fixed",
		},
		VUL_STATUS_REBOOT: ExportHandledVulnerabilitiesRequestBodyStatus{
			value: "vul_status_reboot",
		},
		VUL_STATUS_FAILED: ExportHandledVulnerabilitiesRequestBodyStatus{
			value: "vul_status_failed",
		},
		VUL_STATUS_FIX_AFTER_REBOOT: ExportHandledVulnerabilitiesRequestBodyStatus{
			value: "vul_status_fix_after_reboot",
		},
	}
}

func (c ExportHandledVulnerabilitiesRequestBodyStatus) Value() string {
	return c.value
}

func (c ExportHandledVulnerabilitiesRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHandledVulnerabilitiesRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

type ExportHandledVulnerabilitiesRequestBodyAssetValue struct {
	value string
}

type ExportHandledVulnerabilitiesRequestBodyAssetValueEnum struct {
	IMPORTANT ExportHandledVulnerabilitiesRequestBodyAssetValue
	COMMON    ExportHandledVulnerabilitiesRequestBodyAssetValue
	TEST      ExportHandledVulnerabilitiesRequestBodyAssetValue
}

func GetExportHandledVulnerabilitiesRequestBodyAssetValueEnum() ExportHandledVulnerabilitiesRequestBodyAssetValueEnum {
	return ExportHandledVulnerabilitiesRequestBodyAssetValueEnum{
		IMPORTANT: ExportHandledVulnerabilitiesRequestBodyAssetValue{
			value: "important",
		},
		COMMON: ExportHandledVulnerabilitiesRequestBodyAssetValue{
			value: "common",
		},
		TEST: ExportHandledVulnerabilitiesRequestBodyAssetValue{
			value: "test",
		},
	}
}

func (c ExportHandledVulnerabilitiesRequestBodyAssetValue) Value() string {
	return c.value
}

func (c ExportHandledVulnerabilitiesRequestBodyAssetValue) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHandledVulnerabilitiesRequestBodyAssetValue) UnmarshalJSON(b []byte) error {
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

type ExportHandledVulnerabilitiesRequestBodyType struct {
	value string
}

type ExportHandledVulnerabilitiesRequestBodyTypeEnum struct {
	LINUX_VUL   ExportHandledVulnerabilitiesRequestBodyType
	WINDOWS_VUL ExportHandledVulnerabilitiesRequestBodyType
	WEB_CMS     ExportHandledVulnerabilitiesRequestBodyType
	APP_VUL     ExportHandledVulnerabilitiesRequestBodyType
	URGENT_VUL  ExportHandledVulnerabilitiesRequestBodyType
}

func GetExportHandledVulnerabilitiesRequestBodyTypeEnum() ExportHandledVulnerabilitiesRequestBodyTypeEnum {
	return ExportHandledVulnerabilitiesRequestBodyTypeEnum{
		LINUX_VUL: ExportHandledVulnerabilitiesRequestBodyType{
			value: "linux_vul",
		},
		WINDOWS_VUL: ExportHandledVulnerabilitiesRequestBodyType{
			value: "windows_vul",
		},
		WEB_CMS: ExportHandledVulnerabilitiesRequestBodyType{
			value: "web_cms",
		},
		APP_VUL: ExportHandledVulnerabilitiesRequestBodyType{
			value: "app_vul",
		},
		URGENT_VUL: ExportHandledVulnerabilitiesRequestBodyType{
			value: "urgent_vul",
		},
	}
}

func (c ExportHandledVulnerabilitiesRequestBodyType) Value() string {
	return c.value
}

func (c ExportHandledVulnerabilitiesRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHandledVulnerabilitiesRequestBodyType) UnmarshalJSON(b []byte) error {
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

type ExportHandledVulnerabilitiesRequestBodyHandleCycle struct {
	value string
}

type ExportHandledVulnerabilitiesRequestBodyHandleCycleEnum struct {
	TODAY ExportHandledVulnerabilitiesRequestBodyHandleCycle
	ALL   ExportHandledVulnerabilitiesRequestBodyHandleCycle
}

func GetExportHandledVulnerabilitiesRequestBodyHandleCycleEnum() ExportHandledVulnerabilitiesRequestBodyHandleCycleEnum {
	return ExportHandledVulnerabilitiesRequestBodyHandleCycleEnum{
		TODAY: ExportHandledVulnerabilitiesRequestBodyHandleCycle{
			value: "today",
		},
		ALL: ExportHandledVulnerabilitiesRequestBodyHandleCycle{
			value: "all",
		},
	}
}

func (c ExportHandledVulnerabilitiesRequestBodyHandleCycle) Value() string {
	return c.value
}

func (c ExportHandledVulnerabilitiesRequestBodyHandleCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHandledVulnerabilitiesRequestBodyHandleCycle) UnmarshalJSON(b []byte) error {
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
