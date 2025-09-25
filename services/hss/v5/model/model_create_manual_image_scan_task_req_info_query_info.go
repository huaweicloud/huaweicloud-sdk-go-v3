package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManualImageScanTaskReqInfoQueryInfo 按照查询结果扫描
type CreateManualImageScanTaskReqInfoQueryInfo struct {

	// **参数解释**: 本地镜像类型 **约束限制**: 不涉及 **取值范围**: - swr_image：swr镜像。 - other_image：其他镜像。  **默认取值**: 不涉及
	LocalImageType *string `json:"local_image_type,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位。  **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位。  **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器公网IP **约束限制**: 不涉及 **取值范围**: 字符长度0-128位。  **默认取值**: 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 私有IP地址 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位。  **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 集群id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位。  **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 容器id **约束限制**: 不涉及 **取值范围**: 字符长度0-128位。  **默认取值**: 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 容器名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位。  **默认取值**: 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: pod id **约束限制**: 不涉及 **取值范围**: 字符长度0-64位。  **默认取值**: 不涉及
	PodId *string `json:"pod_id,omitempty"`

	// **参数解释**: pod 名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-512位。  **默认取值**: 不涉及
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**: 软件名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位。  **默认取值**: 不涉及
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 是否存在软件漏洞 **约束限制**: 不涉及 **取值范围**: - true：存在软件漏洞。 - false：不存在软件漏洞。  **默认取值**: 不涉及
	HasVul *bool `json:"has_vul,omitempty"`

	// **参数解释**: 组织名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位。  **默认取值**: 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位。  **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位。  **默认取值**: 不涉及
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - private_image：SWR私有镜像仓库。 - shared_image：SWR共享镜像仓库。 - instance_image：SWR企业仓库。 - harbor：harbor仓库。 - jfrog：jfrog仓库。  **默认取值**: 不涉及
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 扫描状态 **约束限制**: 不涉及 **取值范围**: - unscan：未扫描 - success：扫描完成 - scanning：扫描中 - failed：扫描失败 - download_failed：下载失败 - image_oversized：镜像超大  **默认取值**: 不涉及
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**: 镜像大小 **约束限制**: 不涉及 **取值范围**: 0-9223372036854775807。  **默认取值**: 不涉及
	ImageSize *int64 `json:"image_size,omitempty"`

	// **参数解释**: 创建时间开始日期，时间单位 毫秒（ms） **约束限制**: 不涉及 **取值范围**: 0-9223372036854775807。  **默认取值**: 不涉及
	StartLatestUpdateTime *int64 `json:"start_latest_update_time,omitempty"`

	// **参数解释**: 创建时间结束日期，时间单位 毫秒（ms） **约束限制**: 不涉及 **取值范围**: 0-9223372036854775807。  **默认取值**: 不涉及
	EndLatestUpdateTime *int64 `json:"end_latest_update_time,omitempty"`

	// **参数解释**: 最近一次扫描完成时间开始日期，时间单位 毫秒（ms） **约束限制**: 不涉及 **取值范围**: 0-9223372036854775807。  **默认取值**: 不涉及
	StartLatestScanTime *int64 `json:"start_latest_scan_time,omitempty"`

	// **参数解释**: 最近一次扫描完成时间结束日期，时间单位 毫秒（ms） **约束限制**: 不涉及 **取值范围**: 0-9223372036854775807。  **默认取值**: 不涉及
	EndLatestScanTime *int64 `json:"end_latest_scan_time,omitempty"`
}

func (o CreateManualImageScanTaskReqInfoQueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualImageScanTaskReqInfoQueryInfo struct{}"
	}

	return strings.Join([]string{"CreateManualImageScanTaskReqInfoQueryInfo", string(data)}, " ")
}
