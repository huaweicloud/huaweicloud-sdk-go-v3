package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegistryImagesInfo 镜像信息
type RegistryImagesInfo struct {

	// **参数解释**: id **取值范围**: 最小值0，最大值9223372036854775807
	Id *int64 `json:"id,omitempty"`

	// **参数解释**: 组织名称 **取值范围**: 字符长度0-64位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-128位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像id **取值范围**: 字符长度0-64位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像digest **取值范围**: 字符长度0-128位
	ImageDigest *string `json:"image_digest,omitempty"`

	// **参数解释**: 镜像版本 **取值范围**: 字符长度0-64位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 镜像类型 **取值范围**: - private_image：SWR私有镜像。 - shared_image：SWR共享镜像。 - instance_image：SWR企业版镜像。 - harbor：Harbor仓库镜像。 - jfrog：Jfrog镜像。
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 镜像仓id **取值范围**: 字符长度1-64位
	RegistryId *string `json:"registry_id,omitempty"`

	// **参数解释**: 镜像仓库名称 **取值范围**: 字符长度1-128位
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**： 镜像仓库类型 **取值范围**： - SwrPrivate：swr私有。 - SwrShared：swr共享。 - SwrEnterprise：swr企业。 - Harbor：harbor仓库。 - Jfrog：jfrog仓库。 - Other：其他仓库。
	RegistryType *string `json:"registry_type,omitempty"`

	// 是否是最新版本
	LatestVersion *bool `json:"latest_version,omitempty"`

	// **参数解释**： 扫描状态 **取值范围**： - unscan：未扫描。 - success：扫描完成。 - scanning：正在扫描。 - failed：扫描失败。 - download_failed：下载失败。 - image_oversized：镜像超大。 - waiting_for_scan：等待扫描。
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**： 扫描失败原因描述 **取值范围**： 扫描失败原因code和描述对应关系如下 - unknown_error：未知错误。 - authentication_failed：认证失败。 - download_failed：镜像下载失败。请向技术人员寻求帮助。 - image_over_sized：镜像大小超限，不支持扫描。建议精简镜像。 - get_detail_info_not_found：获取镜像详细信息失败，镜像仓中可能已经不存在此镜像。请重新同步最新镜像。 - image_layer_over_sized：镜像层数超限，不支持扫描。建议精简镜像。 - schema_v1_not_support：Schema V1镜像不支持扫描。建议升级到V2版本。 - access_swr_failed：访问SWR服务出错。请向技术人员寻求帮助。 - swr_authentication_error：缺少SWR授权。请参考镜像授权指导设置权限。 - failed_to_scan_vulnerability：漏洞扫描失败。 - failed_to_scan_file：文件扫描失败。 - failed_to_scan_software：软件扫描失败。 - failed_to_check_sensitive_information：敏感信息核查失败。 - failed_to_check_baseline：基线检查失败。 - failed_to_check_software_compliance：软件合规检查失败。 - failed_to_query_basic_image_information：基础镜像信息查询失败。 - failed_to_check_build_cmd：镜像构建指令扫描失败。 - response_timed_out：响应超时。 - database_error：数据库错误。 - failed_to_send_the_scan_request：发送扫描请求失败。
	ScanFailedDesc *string `json:"scan_failed_desc,omitempty"`

	// **参数解释**： 扫描失败原因code **取值范围**： 扫描失败原因code和描述对应关系如下 - unknown_error：未知错误。 - authentication_failed：认证失败。 - download_failed：镜像下载失败。请向技术人员寻求帮助。 - image_over_sized：镜像大小超限，不支持扫描。建议精简镜像。 - get_detail_info_not_found：获取镜像详细信息失败，镜像仓中可能已经不存在此镜像。请重新同步最新镜像。 - image_layer_over_sized：镜像层数超限，不支持扫描。建议精简镜像。 - schema_v1_not_support：Schema V1镜像不支持扫描。建议升级到V2版本。 - access_swr_failed：访问SWR服务出错。请向技术人员寻求帮助。 - swr_authentication_error：缺少SWR授权。请参考镜像授权指导设置权限。 - failed_to_scan_vulnerability：漏洞扫描失败。 - failed_to_scan_file：文件扫描失败。 - failed_to_scan_software：软件扫描失败。 - failed_to_check_sensitive_information：敏感信息核查失败。 - failed_to_check_baseline：基线检查失败。 - failed_to_check_software_compliance：软件合规检查失败。 - failed_to_query_basic_image_information：基础镜像信息查询失败。 - failed_to_check_build_cmd：镜像构建指令扫描失败。 - response_timed_out：响应超时。 - database_error：数据库错误。 - failed_to_send_the_scan_request：发送扫描请求失败。
	ScanFailedCode *string `json:"scan_failed_code,omitempty"`

	// **参数解释**: 镜像大小 **取值范围**: 取值0-2147483547
	ImageSize *int64 `json:"image_size,omitempty"`

	// **参数解释**: 镜像版本最后更新时间，时间单位 毫秒（ms） **取值范围**: 取值0-9223372036854775807
	LatestUpdateTime *int64 `json:"latest_update_time,omitempty"`

	// **参数解释**: 最近扫描时间，时间单位 毫秒（ms） **取值范围**: 取值0-9223372036854775807
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 最近同步时间，时间单位 毫秒（ms） **取值范围**: 取值0-9223372036854775807
	LatestSyncTime *int64 `json:"latest_sync_time,omitempty"`

	// **参数解释**: 漏洞个数 **取值范围**: 取值0-2147483647
	VulNum *int32 `json:"vul_num,omitempty"`

	// **参数解释**: 基线扫描未通过数 **取值范围**: 取值0-2147483647
	UnsafeSettingNum *int32 `json:"unsafe_setting_num,omitempty"`

	// **参数解释**: 恶意文件数 **取值范围**: 取值0-2147483647
	MaliciousFileNum *int32 `json:"malicious_file_num,omitempty"`

	// **参数解释**: 拥有者（共享镜像参数） **取值范围**: 字符长度0-128
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**： 共享镜像状态 **取值范围**： - expired：已过期。 - effective：有效。
	SharedStatus *string `json:"shared_status,omitempty"`

	// 是否可扫描
	Scannable *bool `json:"scannable,omitempty"`

	// 是否是多架构镜像
	IsMultipleArch *bool `json:"is_multiple_arch,omitempty"`

	// **参数解释**: 企业版镜像实例名称 **取值范围**: 字符长度0-128
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 企业版镜像实例ID **取值范围**: 字符长度0-64
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 企业版镜像实例URL **取值范围**: 字符长度0-256
	InstanceUrl *string `json:"instance_url,omitempty"`

	// **参数解释**: 镜像风险程度，在镜像扫描完成后展示 **取值范围**: - Security：安全。 - Low：低危。 - Medium：中危。 - High：高危。
	SeverityLevel *string `json:"severity_level,omitempty"`

	// 多架构关联镜像信息
	AssociationImages *[]AssociateImagesInfo `json:"association_images,omitempty"`
}

func (o RegistryImagesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegistryImagesInfo struct{}"
	}

	return strings.Join([]string{"RegistryImagesInfo", string(data)}, " ")
}
