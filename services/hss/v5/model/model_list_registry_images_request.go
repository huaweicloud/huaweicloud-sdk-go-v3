package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegistryImagesRequest Request Object
type ListRegistryImagesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 组织名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位。  **默认取值**: 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 仓库名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - private_image：SWR私有镜像。 - shared_image：SWR共享镜像。 - instance_image：SWR企业版镜像。 - harbor：Harbor仓库镜像。 - jfrog：Jfrog镜像。  **默认取值**: 不涉及
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 可排序字段 **约束限制**: 不涉及 **取值范围**: - latest_scan_time：最近扫描时间。  **默认取值**: 不涉及
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 排序的顺序 **约束限制**: 不涉及 **取值范围**:   - asc  : 正序   - desc : 倒序  **默认取值**: 正序排序
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值1-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 仅关注最新版本镜像 **约束限制**: 不涉及 **取值范围**: - true：是。 - false：否。  **默认取值**: false
	LatestVersion *bool `json:"latest_version,omitempty"`

	// **参数解释**: 镜像大小 **约束限制**: 取值为非负整数，单位为字节（bytes） **取值范围**: 取值0-2147483547 **默认取值**: 不涉及
	ImageSize *int64 `json:"image_size,omitempty"`

	// **参数解释**: 扫描状态 **约束限制**: 不涉及 **取值范围**: - unscan：未扫描。 - success：扫描完成。 - scanning：扫描中。 - failed：扫描失败。 - waiting_for_scan：等待扫描。  **默认取值**: 不涉及
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**: 最后更新时间的查询起始值，时间单位 毫秒（ms） **约束限制**: 需≤end_latest_update_time，未传end参数时默认查询至当前时间 **取值范围**: 取值0-9223372036854775807 ms（UTC时区，1970-01-01 00:00:00起） **默认取值**: 不涉及
	StartLatestUpdateTime *int64 `json:"start_latest_update_time,omitempty"`

	// **参数解释**: 最后更新时间的查询结束值，时间单位 毫秒（ms） **约束限制**: 需≥start_latest_update_time，未传start参数时默认从0开始 **取值范围**: 取值0-9223372036854775807 ms（UTC时区，1970-01-01 00:00:00起） **默认取值**: 不涉及
	EndLatestUpdateTime *int64 `json:"end_latest_update_time,omitempty"`

	// **参数解释**: 最近一次扫描完成时间的查询起始值，时间单位 毫秒（ms） **约束限制**: 需≤end_latest_scan_time，未传end参数时默认查询至当前时间 **取值范围**: 取值0-9223372036854775807 ms（UTC时区，1970-01-01 00:00:00起） **默认取值**: 不涉及
	StartLatestScanTime *int64 `json:"start_latest_scan_time,omitempty"`

	// **参数解释**: 最近一次扫描完成时间的查询结束值，时间单位 毫秒（ms） **约束限制**: 需≥start_latest_scan_time，未传start参数时默认从0开始 **取值范围**: 取值0-9223372036854775807 ms（UTC时区，1970-01-01 00:00:00起） **默认取值**: 不涉及
	EndLatestScanTime *int64 `json:"end_latest_scan_time,omitempty"`

	// **参数解释**: 最近一次同步完成时间的查询起始值，时间单位 毫秒（ms） **约束限制**: 需≤end_latest_sync_time，未传end参数时默认查询至当前时间 **取值范围**: 取值0-9223372036854775807 ms（UTC时区，1970-01-01 00:00:00起） **默认取值**: 不涉及
	StartLatestSyncTime *int64 `json:"start_latest_sync_time,omitempty"`

	// **参数解释**: 最近一次同步完成时间的查询结束值，时间单位 毫秒（ms） **约束限制**: 需≥start_latest_sync_time，未传start参数时默认从0开始 **取值范围**: 取值0-9223372036854775807 ms（UTC时区，1970-01-01 00:00:00起） **默认取值**: 不涉及
	EndLatestSyncTime *int64 `json:"end_latest_sync_time,omitempty"`

	// **参数解释**: 是否存在恶意文件 **约束限制**: 不涉及 **取值范围**: - true：是。 - false：否。  **默认取值**: 不涉及
	HasMaliciousFile *bool `json:"has_malicious_file,omitempty"`

	// **参数解释**： 是否存在基线检查风险 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	HasUnsafeSetting *bool `json:"has_unsafe_setting,omitempty"`

	// **参数解释**： 是否存在软件漏洞 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	HasVul *bool `json:"has_vul,omitempty"`

	// **参数解释**: 有安全风险 **约束限制**: 不涉及 **取值范围**: - true：是。 - false：否。  **默认取值**: 不涉及
	Risky *bool `json:"risky,omitempty"`

	// **参数解释**： 企业仓库实例ID，SWR企业版可以使用该参数，需要到SWR企业版服务中获取 **约束限制**： 仅SWR企业版镜像查询有效 **取值范围**： 字符长度0-128位 **默认取值**： 不涉及
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 企业仓库实例名称，SWR企业版可以使用该参数，需要到SWR企业版服务中获取 **约束限制**： 仅SWR企业版镜像查询有效 **取值范围**： 字符长度0-128位 **默认取值**： 不涉及
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 是否是多架构镜像 **约束限制**: 不涉及 **取值范围**: - true：是。 - false：否。  **默认取值**: 不涉及
	IsMultarch *bool `json:"is_multarch,omitempty"`

	// **参数解释**: 镜像风险程度，在镜像扫描完成后展示 **约束限制**: 不涉及 **取值范围**: - Security：安全。 - Low：低危。 - Medium：中危。 - High：高危。  **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`
}

func (o ListRegistryImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegistryImagesRequest struct{}"
	}

	return strings.Join([]string{"ListRegistryImagesRequest", string(data)}, " ")
}
