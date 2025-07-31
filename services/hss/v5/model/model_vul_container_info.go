package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulContainerInfo 软件漏洞列表
type VulContainerInfo struct {

	// 受漏洞影响的容器id
	ContainerId *string `json:"container_id,omitempty"`

	// 受影响容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 危险程度   - Critical : 漏洞cvss评分大于等于9；对应控制台页面的高危   - High : 漏洞cvss评分大于等于7，小于9；对应控制台页面的中危   - Medium : 漏洞cvss评分大于等于4，小于7；对应控制台页面的中危   - Low : 漏洞cvss评分小于4；对应控制台页面的低危
	SeverityLevel *string `json:"severity_level,omitempty"`

	// 受漏洞影响的容器所处集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 受影响容器所处集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 受漏洞影响的容器对应镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 受影响容器对应镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 镜像类型
	ImageType *string `json:"image_type,omitempty"`

	// 所属组织
	ImageOrg *string `json:"image_org,omitempty"`

	// 仓库类型
	RegistryType *string `json:"registry_type,omitempty"`

	// 仓库名称
	RegistryName *string `json:"registry_name,omitempty"`

	// 受影响容器对应主机名称
	HostName *string `json:"host_name,omitempty"`

	// 漏洞状态   - vul_status_unfix : 未处理   - vul_status_ignored : 已忽略   - vul_status_verified : 验证中   - vul_status_fixing : 修复中   - vul_status_fixed : 修复成功   - vul_status_reboot : 修复成功待重启   - vul_status_failed : 修复失败   - vul_status_fix_after_reboot : 请重启主机再次修复
	Status *string `json:"status,omitempty"`

	// 服务器公网ip
	HostIp *string `json:"host_ip,omitempty"`

	// 服务器私网ip
	PrivateIp *string `json:"private_ip,omitempty"`

	// 处置操作的描述信息
	Remark *string `json:"remark,omitempty"`

	// 首次扫描时间
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// 扫描时间，时间戳单位：毫秒
	ScanTime *int64 `json:"scan_time,omitempty"`
}

func (o VulContainerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulContainerInfo struct{}"
	}

	return strings.Join([]string{"VulContainerInfo", string(data)}, " ")
}
