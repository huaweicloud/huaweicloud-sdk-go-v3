package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportVulHandleHistoryRequestBody struct {

	// 漏洞状态 vul_status_unfix : 未处理 vul_status_ignored : 已忽略 vul_status_verified : 验证中 vul_status_fixing : 修复中 vul_status_fixed : 修复成功 vul_status_reboot : 修复成功待重启 vul_status_failed : 修复失败 vul_status_fix_after_reboot : 请重启主机再次修复
	Status *ExportVulHandleHistoryRequestBodyStatus `json:"status,omitempty"`

	// 漏洞ID
	VulId *string `json:"vul_id,omitempty"`

	// 漏洞类型，包含如下：   -linux_vul : linux漏洞   -windows_vul : windows漏洞   -web_cms : Web-CMS漏洞   -app_vul : 应用漏洞   -urgent_vul : 应急漏洞
	VulType *ExportVulHandleHistoryRequestBodyVulType `json:"vul_type,omitempty"`

	// 资产重要性，包含如下3种 important ：重要资产 common ：一般资产 test ：测试资产
	AssetValue *ExportVulHandleHistoryRequestBodyAssetValue `json:"asset_value,omitempty"`

	// 服务器组
	GroupName *string `json:"group_name,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ip
	HostIp *string `json:"host_ip,omitempty"`

	// 导出数据条数
	ExportSize *int32 `json:"export_size,omitempty"`

	// 导出漏洞数据的表头信息列表
	ExportHeaderList *[][]string `json:"export_header_list,omitempty"`
}

func (o ExportVulHandleHistoryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVulHandleHistoryRequestBody struct{}"
	}

	return strings.Join([]string{"ExportVulHandleHistoryRequestBody", string(data)}, " ")
}

type ExportVulHandleHistoryRequestBodyStatus struct {
	value string
}

type ExportVulHandleHistoryRequestBodyStatusEnum struct {
	VUL_STATUS_UNFIX            ExportVulHandleHistoryRequestBodyStatus
	VUL_STATUS_IGNORED          ExportVulHandleHistoryRequestBodyStatus
	VUL_STATUS_VERIFIED         ExportVulHandleHistoryRequestBodyStatus
	VUL_STATUS_FIXING           ExportVulHandleHistoryRequestBodyStatus
	VUL_STATUS_FIXED            ExportVulHandleHistoryRequestBodyStatus
	VUL_STATUS_REBOOT           ExportVulHandleHistoryRequestBodyStatus
	VUL_STATUS_FAILED           ExportVulHandleHistoryRequestBodyStatus
	VUL_STATUS_FIX_AFTER_REBOOT ExportVulHandleHistoryRequestBodyStatus
}

func GetExportVulHandleHistoryRequestBodyStatusEnum() ExportVulHandleHistoryRequestBodyStatusEnum {
	return ExportVulHandleHistoryRequestBodyStatusEnum{
		VUL_STATUS_UNFIX: ExportVulHandleHistoryRequestBodyStatus{
			value: "vul_status_unfix",
		},
		VUL_STATUS_IGNORED: ExportVulHandleHistoryRequestBodyStatus{
			value: "vul_status_ignored",
		},
		VUL_STATUS_VERIFIED: ExportVulHandleHistoryRequestBodyStatus{
			value: "vul_status_verified",
		},
		VUL_STATUS_FIXING: ExportVulHandleHistoryRequestBodyStatus{
			value: "vul_status_fixing",
		},
		VUL_STATUS_FIXED: ExportVulHandleHistoryRequestBodyStatus{
			value: "vul_status_fixed",
		},
		VUL_STATUS_REBOOT: ExportVulHandleHistoryRequestBodyStatus{
			value: "vul_status_reboot",
		},
		VUL_STATUS_FAILED: ExportVulHandleHistoryRequestBodyStatus{
			value: "vul_status_failed",
		},
		VUL_STATUS_FIX_AFTER_REBOOT: ExportVulHandleHistoryRequestBodyStatus{
			value: "vul_status_fix_after_reboot",
		},
	}
}

func (c ExportVulHandleHistoryRequestBodyStatus) Value() string {
	return c.value
}

func (c ExportVulHandleHistoryRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportVulHandleHistoryRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

type ExportVulHandleHistoryRequestBodyVulType struct {
	value string
}

type ExportVulHandleHistoryRequestBodyVulTypeEnum struct {
	LINUX_VUL   ExportVulHandleHistoryRequestBodyVulType
	WINDOWS_VUL ExportVulHandleHistoryRequestBodyVulType
	WEB_CMS     ExportVulHandleHistoryRequestBodyVulType
	APP_VUL     ExportVulHandleHistoryRequestBodyVulType
	URGENT_VUL  ExportVulHandleHistoryRequestBodyVulType
}

func GetExportVulHandleHistoryRequestBodyVulTypeEnum() ExportVulHandleHistoryRequestBodyVulTypeEnum {
	return ExportVulHandleHistoryRequestBodyVulTypeEnum{
		LINUX_VUL: ExportVulHandleHistoryRequestBodyVulType{
			value: "linux_vul",
		},
		WINDOWS_VUL: ExportVulHandleHistoryRequestBodyVulType{
			value: "windows_vul",
		},
		WEB_CMS: ExportVulHandleHistoryRequestBodyVulType{
			value: "web_cms",
		},
		APP_VUL: ExportVulHandleHistoryRequestBodyVulType{
			value: "app_vul",
		},
		URGENT_VUL: ExportVulHandleHistoryRequestBodyVulType{
			value: "urgent_vul",
		},
	}
}

func (c ExportVulHandleHistoryRequestBodyVulType) Value() string {
	return c.value
}

func (c ExportVulHandleHistoryRequestBodyVulType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportVulHandleHistoryRequestBodyVulType) UnmarshalJSON(b []byte) error {
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

type ExportVulHandleHistoryRequestBodyAssetValue struct {
	value string
}

type ExportVulHandleHistoryRequestBodyAssetValueEnum struct {
	IMPORTANT ExportVulHandleHistoryRequestBodyAssetValue
	COMMON    ExportVulHandleHistoryRequestBodyAssetValue
	TEST      ExportVulHandleHistoryRequestBodyAssetValue
}

func GetExportVulHandleHistoryRequestBodyAssetValueEnum() ExportVulHandleHistoryRequestBodyAssetValueEnum {
	return ExportVulHandleHistoryRequestBodyAssetValueEnum{
		IMPORTANT: ExportVulHandleHistoryRequestBodyAssetValue{
			value: "important",
		},
		COMMON: ExportVulHandleHistoryRequestBodyAssetValue{
			value: "common",
		},
		TEST: ExportVulHandleHistoryRequestBodyAssetValue{
			value: "test",
		},
	}
}

func (c ExportVulHandleHistoryRequestBodyAssetValue) Value() string {
	return c.value
}

func (c ExportVulHandleHistoryRequestBodyAssetValue) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportVulHandleHistoryRequestBodyAssetValue) UnmarshalJSON(b []byte) error {
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
