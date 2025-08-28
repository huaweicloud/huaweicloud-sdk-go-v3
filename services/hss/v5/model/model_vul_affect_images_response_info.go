package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulAffectImagesResponseInfo **参数解释**: 受漏洞影响的镜像列表 **取值范围**: 不涉及
type VulAffectImagesResponseInfo struct {

	// **参数解释**: 镜像id **取值范围**: 字符长度0-256位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-512位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 漏洞id **取值范围**: 字符长度0-256位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 危险程度 **取值范围**: - Low：低危 - Medium：中危 - High：高危 - Critical：严重
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释**: 关联容器数 **取值范围**: 取值0-2147483547
	ContainerNum *int32 `json:"container_num,omitempty"`

	// **参数解释**: 镜像标识 **取值范围**: 字符长度0-256位
	ImageDigest *string `json:"image_digest,omitempty"`

	// **参数解释**: 镜像版本 **取值范围**: 字符长度0-64位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: Agent ID **取值范围**: 字符长度0-128位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 漏洞当前状态 **取值范围**: - vul_status_unfix：未处理 - vul_status_ignored：已忽略
	Status *string `json:"status,omitempty"`

	// **参数解释**: 首次扫描时间（时间戳，单位为毫秒） **取值范围**: 字符长度0-32位
	FirstScanTime *string `json:"first_scan_time,omitempty"`

	// **参数解释**: 最近扫描时间（时间戳，单位为毫秒） **取值范围**: 字符长度0-32位
	LatestScanTime *string `json:"latest_scan_time,omitempty"`

	// **参数解释**: 镜像类型 **取值范围**: - local_image：本地镜像 - registry：仓库镜像 - cicd：CI/CD镜像
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 镜像大小 **取值范围**: 字符长度0-128位
	ImageSize *string `json:"image_size,omitempty"`

	// **参数解释**: 所属组织 **取值范围**: 字符长度0-256位
	Organization *string `json:"organization,omitempty"`

	// **参数解释**: 镜像仓类型 **取值范围**: - Harbor：harbor - Jfrog：jfrog - SwrPrivate：swr私有 - SwrShared：swr共享 - SwrEnterprise：swr企业 - Other：其他仓库
	RegistryType *string `json:"registry_type,omitempty"`

	// **参数解释**: 仓库名称 **取值范围**: 字符长度0-256位
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**: 是否为多架构 **取值范围**: - true：是多架构 - false：不是多架构
	IsMultiArch *string `json:"is_multi_arch,omitempty"`

	// **参数解释**: 集群ID **取值范围**: 字符长度0-128位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 集群名称 **取值范围**: 字符长度0-256位
	ClusterName *string `json:"cluster_name,omitempty"`
}

func (o VulAffectImagesResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectImagesResponseInfo struct{}"
	}

	return strings.Join([]string{"VulAffectImagesResponseInfo", string(data)}, " ")
}
