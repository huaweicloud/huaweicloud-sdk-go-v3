package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkLoadInfo 工作负载信息
type WorkLoadInfo struct {

	// **参数解释**: 所属企业项目名称 **取值范围**: 字符长度0-65535位
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释**: 工作负载名称 **取值范围**: 字符长度0-65535位
	WorkloadName *string `json:"workload_name,omitempty"`

	// **参数解释**: 工作负载Id **取值范围**: 字符长度0-65535位
	WorkloadId *string `json:"workload_id,omitempty"`

	// **参数解释**: 工作负载类型 **取值范围**: 字符长度0-65535位
	WorkloadType *string `json:"workload_type,omitempty"`

	// **参数解释**: 工作负载状态 **取值范围**: 字符长度0-65535位
	WorkloadStatus *string `json:"workload_status,omitempty"`

	// **参数解释**: 命名空间名称 **取值范围**: 字符长度0-256位
	NamespaceName *string `json:"namespace_name,omitempty"`

	// **参数解释**: 所属集群名称 **取值范围**: 字符长度0-65535位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**: 所属集群ID **取值范围**: 字符长度0-65535位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: agent版本 **取值范围**: 字符长度0-65535位
	AgentVersion *string `json:"agent_version,omitempty"`

	// **参数解释**: 已安装实例个数 **取值范围**: 最小值0，最大值2147483647
	InstalledPodNum *int32 `json:"installed_pod_num,omitempty"`

	// **参数解释**: 未安装实例个数 **取值范围**: 最小值0，最大值2147483647
	NotInstallNum *int32 `json:"not_install_num,omitempty"`

	// **参数解释**: 安装失败实例个数 **取值范围**: 最小值0，最大值2147483647
	InstallFailedNum *int32 `json:"install_failed_num,omitempty"`

	// **参数解释**: 离线agent个数 **取值范围**: 最小值0，最大值2147483647
	OfflineNum *int32 `json:"offline_num,omitempty"`

	// **参数解释**: 全部实例个数 **取值范围**: 最小值0，最大值2147483647
	PodNum *int32 `json:"pod_num,omitempty"`

	// **参数解释**: 已防护实例个数 **取值范围**: 最小值0，最大值2147483647
	ProtectPodNum *int32 `json:"protect_pod_num,omitempty"`

	// **参数解释**: 运行任务数 **取值范围**: 最小值0，最大值2147483647
	RunningJobNum *int32 `json:"running_job_num,omitempty"`

	// **参数解释**: 镜像名称列表 **取值范围**: 最小值0，最大值256
	ImageName *[]string `json:"image_name,omitempty"`

	// **参数解释**: 标签 **取值范围**: 最小值0，最大值65535
	Labels *[]LabelInfo `json:"labels,omitempty"`

	// **参数解释**: 实例信息 **取值范围**: 最小值0，最大值65535
	InstanceInfo *[]ServerlessInstanceInfo `json:"instance_info,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: 字符长度0-65535位
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: 普通任务为完成时间，定时任务类型为最近调度时间 **取值范围**: 最小值0，最大值9223372036854775807
	LastScheduleTime *int64 `json:"last_schedule_time,omitempty"`

	// **参数解释**: 定时任务类型 任务触发 **取值范围**: 字符长度0-65535位
	Schedule *string `json:"schedule,omitempty"`

	// **参数解释**: 是否可升级 **取值范围**: - true：是。 - false：否。
	Upgradable *bool `json:"upgradable,omitempty"`
}

func (o WorkLoadInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkLoadInfo struct{}"
	}

	return strings.Join([]string{"WorkLoadInfo", string(data)}, " ")
}
