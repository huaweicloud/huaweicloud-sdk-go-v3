package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExportSwrBaselineInfoRequestInfo 操作的事件
type BatchExportSwrBaselineInfoRequestInfo struct {

	// **参数解释**: 要导出的镜像信息列表，operate_all参数为false时需要填写批量查询条件,image_id 镜像id，唯一镜像标识（注：对私有镜像和共享镜像来说是镜像列表返回的id） **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2147483647 **默认取值**: 不涉及
	ImageIdList *[]string `json:"image_id_list,omitempty"`

	// **参数解释**： 若为true全量查询，可筛选条件全部查询 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	OperateAll *bool `json:"operate_all,omitempty"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - private_image: 私有镜像。 - shared_image: 共享镜像。 - local_image: 本地镜像。 - instance_image: 企业镜像。 - harbor: Harbor镜像。 - jfrog: Jfrog镜像。 - cicd: cicd镜像。  **默认取值**: 不涉及
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: - linux_vul : linux漏洞。 - app_vul : 应用漏洞。  **默认取值**: 不涉及
	VulType *string `json:"vul_type,omitempty"`

	// **参数解释**: 扫描状态 **约束限制**: 不涉及 **取值范围**: - unscan: 未扫描。 - success: 扫描完成。 - scanning: 扫描中。 - failed: 扫描失败。 - download_failed: 下载失败。 - image_oversized: 镜像超大。  **默认取值**: 不涉及
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**: 组织名称（只有私有镜像和共享镜像有该字段，本地镜像没有） **约束限制**: 不涉及 **取值范围**: 字符长度0-65535位 **默认取值**: 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-65535位 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-65535位 **默认取值**: 不涉及
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 镜像id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 仅关注最新版本镜像 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	LatestVersion *bool `json:"latest_version,omitempty"`

	// **参数解释**: 镜像大小 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2147483547 **默认取值**: 不涉及
	ImageSize *int64 `json:"image_size,omitempty"`

	// **参数解释**: 创建时间开始日期，时间单位 毫秒（ms） **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2147483547 **默认取值**: 不涉及
	StartLatestUpdateTime *int64 `json:"start_latest_update_time,omitempty"`

	// **参数解释**: 创建时间结束日期，时间单位 毫秒（ms） **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2147483547 **默认取值**: 不涉及
	EndLatestUpdateTime *int64 `json:"end_latest_update_time,omitempty"`

	// **参数解释**: 最近一次扫描完成时间开始日期，时间单位 毫秒（ms） **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2147483547 **默认取值**: 不涉及
	StartLatestScanTime *int64 `json:"start_latest_scan_time,omitempty"`

	// **参数解释**: 最近一次扫描完成时间结束日期，时间单位 毫秒（ms） **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2147483547 **默认取值**: 不涉及
	EndLatestScanTime *int64 `json:"end_latest_scan_time,omitempty"`

	// **参数解释**： 是否存在恶意文件 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	HasMaliciousFile *bool `json:"has_malicious_file,omitempty"`

	// **参数解释**： 是否存在软件漏洞 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	HasVul *bool `json:"has_vul,omitempty"`

	// **参数解释**： 是否存在基线检查 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	HasUnsafeSetting *bool `json:"has_unsafe_setting,omitempty"`

	// **参数解释**： 是否有安全风险 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	Risky *bool `json:"risky,omitempty"`

	// **参数解释**: 镜像风险程度，在镜像扫描完成后展示，包含如下: **约束限制**: 不涉及 **取值范围**: - Security : 安全。 - Low : 低危。 - Medium : 中危。 - High : 高危。  **默认取值**: 不涉及
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释**: 企业镜像实例名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 企业仓库实例ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: cicd名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	CicdName *string `json:"cicd_name,omitempty"`

	// **参数解释**: 构建指令风险程度 **约束限制**: 不涉及 **取值范围**: - critical ：严重。 - high ：高危。 - medium ：中危。 - low ：低危。  **默认取值**: 不涉及
	BuildCommandRiskLevel *string `json:"build_command_risk_level,omitempty"`

	// **参数解释**: 构建指令风险名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	BuildCommandRiskName *string `json:"build_command_risk_name,omitempty"`

	// **参数解释**: 要导出的构建指令风险检测规则id列表。rule_id的值可以从接口/v5/{project_id}/image/build-command-risks的返回参数获取 **约束限制**: 不涉及 **取值范围**: 字符长度1-200位 **默认取值**: 不涉及
	BuildCommandRuleIdList *[]string `json:"build_command_rule_id_list,omitempty"`

	// **参数解释**： 是否存在容器 **约束限制**: 不涉及 **取值范围**： - true：是。 - false：否。  **默认取值**: 不涉及
	HasContainer *bool `json:"has_container,omitempty"`

	// **参数解释**: 导出镜像漏洞时的漏洞id列表 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2048 **默认取值**: 不涉及
	VulIdList *[]string `json:"vul_id_list,omitempty"`
}

func (o BatchExportSwrBaselineInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportSwrBaselineInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"BatchExportSwrBaselineInfoRequestInfo", string(data)}, " ")
}
