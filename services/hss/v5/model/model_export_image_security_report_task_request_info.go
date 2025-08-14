package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportImageSecurityReportTaskRequestInfo 镜像安全报告导出任务请求信息。
type ExportImageSecurityReportTaskRequestInfo struct {

	// 导出数据条数
	ExportSize int32 `json:"export_size"`

	// **参数解释**: 安全报告类型 **约束限制**: 不涉及 **取值范围**: - malwares : 恶意文件信息 - apps : 软件信息 - files : 文件信息 - sensitive-info : 敏感信息 - non-compliant-app : 软件合规信息 - build-command-risks : 镜像构建指令风险
	SecurityReportType string `json:"security_report_type"`

	// 要导出的镜像id列表。operate_all参数为false时,需要填写此批量查询条件。本地/CICD镜像填写image_id，仓库镜像填写镜像列表返回的id
	ImageIdList *[]string `json:"image_id_list,omitempty"`

	// 若为true表示查询全量镜像，可传其他参数对全量镜像数据进行筛选
	OperateAll bool `json:"operate_all"`

	// **参数解释**: 镜像类型。导出本地镜像和cicd镜像报告时，需要传参image_type。不传此参数时，默认导出其他仓库镜像数据。 **约束限制**: 不涉及 **取值范围**: - private_image : 私有镜像 - shared_image : 共享镜像 - local_image : 本地镜像 - instance_image : 企业镜像 - harbor : Harbor镜像 - jfrog : Jfrog镜像 - cicd : cicd镜像
	ImageType *string `json:"image_type,omitempty"`

	// 扫描状态，包含如下:   - unscan : 未扫描   - success : 扫描完成   - scanning : 扫描中   - failed : 扫描失败   - download_failed : 下载失败   - image_oversized : 镜像超大
	ScanStatus *string `json:"scan_status,omitempty"`

	// 组织名称
	Namespace *string `json:"namespace,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 仅关注最新版本镜像
	LatestVersion *bool `json:"latest_version,omitempty"`

	// 镜像大小
	ImageSize *int64 `json:"image_size,omitempty"`

	// 创建时间开始日期，时间单位：毫秒（ms）
	StartLatestUpdateTime *int64 `json:"start_latest_update_time,omitempty"`

	// 创建时间结束日期，时间单位：毫秒（ms）
	EndLatestUpdateTime *int64 `json:"end_latest_update_time,omitempty"`

	// 最近一次扫描完成时间开始日期，时间单位：毫秒（ms）
	StartLatestScanTime *int64 `json:"start_latest_scan_time,omitempty"`

	// 最近一次扫描完成时间结束日期，时间单位：毫秒（ms）
	EndLatestScanTime *int64 `json:"end_latest_scan_time,omitempty"`

	// 是否存在恶意文件
	HasMaliciousFile *bool `json:"has_malicious_file,omitempty"`

	// 是否存在软件漏洞
	HasVul *bool `json:"has_vul,omitempty"`

	// 是否存在基线检查
	HasUnsafeSetting *bool `json:"has_unsafe_setting,omitempty"`

	// 是否有安全风险，包含如下:   - true : 有风险   - false : 无风险
	Risky *bool `json:"risky,omitempty"`

	// 镜像风险程度，在镜像扫描完成后展示，包含如下：   - Security：安全   - Low：低危   - Medium：中危   - High：高危
	SeverityLevel *string `json:"severity_level,omitempty"`

	// 企业镜像实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// cicd名称
	CicdName *string `json:"cicd_name,omitempty"`

	// **参数解释**: 构建指令风险程度 **约束限制**: 不涉及 **取值范围**: - critical ：严重 - high ：高危 - medium ：中危 - low ：低危  **默认取值**: 不涉及
	BuildCommandRiskLevel *string `json:"build_command_risk_level,omitempty"`

	// 构建指令风险名称
	BuildCommandRiskName *string `json:"build_command_risk_name,omitempty"`

	// 要导出的构建指令风险检测规则id列表。rule_id的值可以从接口/v5/{project_id}/image/build-command-risks的返回参数获取
	BuildCommandRuleIdList *[]string `json:"build_command_rule_id_list,omitempty"`
}

func (o ExportImageSecurityReportTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageSecurityReportTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"ExportImageSecurityReportTaskRequestInfo", string(data)}, " ")
}
