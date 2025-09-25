package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CicdImageResponseInfo cicd image info
type CicdImageResponseInfo struct {

	// **参数解释**: cicd镜像标识 **取值范围**: 字符长度0-128位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: cicd名称 **取值范围**: 字符长度0-128位
	CicdName *string `json:"cicd_name,omitempty"`

	// **参数解释**: 命名空间 **取值范围**: 字符长度0-64位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-128位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像digest **取值范围**: 字符长度0-128位
	ImageDigest *string `json:"image_digest,omitempty"`

	// **参数解释**: 镜像版本 **取值范围**: 字符长度0-64位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 镜像仓库名称 **取值范围**: 字符长度1-128位
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**： 镜像类型 **取值范围**： - cicd ：cicd镜像。
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**： 镜像仓库类型 **取值范围**： - cicd ：cicd镜像。
	RegistryType *string `json:"registry_type,omitempty"`

	// **参数解释**: 是否是最新版本 **取值范围**: - true：是。 - false：否。
	LatestVersion *bool `json:"latest_version,omitempty"`

	// **参数解释**: 扫描状态 **取值范围**: - unscan ：未扫描。 - success ：扫描完成。 - scanning ：正在扫描。 - failed ：扫描失败。 - download_failed ：下载失败。 - image_oversized ：镜像超大。 - waiting_for_scan ：等待扫描。
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**: 扫描失败原因，包含如下16种。   - 未知错误   - 认证失败   - 镜像下载失败。请向技术人员寻求帮助。   - 镜像大小超限，不支持扫描。建议精简镜像。   - 获取镜像详细信息失败，镜像仓中可能已经不存在此镜像。请重新同步最新镜像。   - 镜像层数超限，不支持扫描。建议精简镜像。   - 漏洞扫描失败   - 文件扫描失败   - 软件扫描失败   - 敏感信息核查失败   - 基线检查失败   - 软件合规检查失败   - 基础镜像信息查询失败   - 响应超时   - 数据库错误   - 发送扫描请求失败 **取值范围**: 字符长度0-128位
	ScanFailedDesc *string `json:"scan_failed_desc,omitempty"`

	// **参数解释**: 扫描失败原因code，包含如下16种。   - \"unknown_error\"   - \"authentication_failed\"   - \"download_failed\"   - \"image_over_sized\"   - \"get_detail_info_not_found\"   - \"image_layer_over_sized\"   - \"failed_to_scan_vulnerability\"   - \"failed_to_scan_file\"   - \"failed_to_scan_software\"   - \"failed_to_check_sensitive_information\"   - \"failed_to_check_baseline\"   - \"failed_to_check_software_compliance\"   - \"failed_to_query_basic_image_information\"   - \"response_timed_out\"   - \"database_error\"   - \"failed_to_send_the_scan_request\" **取值范围**: 字符长度0-64位
	ScanFailedCode *string `json:"scan_failed_code,omitempty"`

	// **参数解释**: 镜像大小 **取值范围**: 最小值0，最大值2147483547
	ImageSize *int64 `json:"image_size,omitempty"`

	// **参数解释**: 镜像版本最后更新时间，时间单位 毫秒（ms） **取值范围**: 最小值0，最大值9223372036854775807
	LatestUpdateTime *int64 `json:"latest_update_time,omitempty"`

	// **参数解释**: 最近扫描时间，时间单位 毫秒（ms） **取值范围**: 最小值0，最大值9223372036854775807
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 漏洞个数 **取值范围**: 最小值0，最大值2147483647
	VulNum *int32 `json:"vul_num,omitempty"`

	// **参数解释**: 基线扫描未通过数 **取值范围**: 最小值0，最大值2147483647
	UnsafeSettingNum *int32 `json:"unsafe_setting_num,omitempty"`

	// **参数解释**: 恶意文件数 **取值范围**: 最小值0，最大值2147483647
	MaliciousFileNum *int32 `json:"malicious_file_num,omitempty"`

	// **参数解释**: 镜像风险程度，在镜像扫描完成后展示 **取值范围**: - Security：安全。 - Low：低危。 - Medium：中危。 - High：高危。
	SeverityLevel *string `json:"severity_level,omitempty"`
}

func (o CicdImageResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CicdImageResponseInfo struct{}"
	}

	return strings.Join([]string{"CicdImageResponseInfo", string(data)}, " ")
}
