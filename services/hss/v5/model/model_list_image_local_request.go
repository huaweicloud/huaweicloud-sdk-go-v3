package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageLocalRequest Request Object
type ListImageLocalRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 扫描状态，包含如下:   - unscan : 未扫描   - success : 扫描完成   - scanning : 扫描中   - failed : 扫描失败   - waiting_for_scan : 等待扫描
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释** : 镜像类型 **约束限制** : 不涉及 **取值范围** : 包含如下两种:  - other_image : 非SWR镜像  - swr_image : SWR镜像 **默认取值** : 不涉及
	LocalImageType *string `json:"local_image_type,omitempty"`

	// 镜像大小，单位字节
	ImageSize *int64 `json:"image_size,omitempty"`

	// 最近更新时间搜索开始日期，时间单位：毫秒（ms）
	StartLatestUpdateTime *int64 `json:"start_latest_update_time,omitempty"`

	// 最近更新时间搜索结束日期，时间单位：毫秒（ms）
	EndLatestUpdateTime *int64 `json:"end_latest_update_time,omitempty"`

	// 最近一次扫描完成时间搜索开始日期，时间单位：毫秒（ms）
	StartLatestScanTime *int64 `json:"start_latest_scan_time,omitempty"`

	// 最近一次扫描完成时间搜索结束日期，时间单位：毫秒（ms）
	EndLatestScanTime *int64 `json:"end_latest_scan_time,omitempty"`

	// 是否存在软件漏洞
	HasVul *bool `json:"has_vul,omitempty"`

	// 本地镜像所关联服务器的名称
	HostName *string `json:"host_name,omitempty"`

	// 本地镜像所关联服务器的ID
	HostId *string `json:"host_id,omitempty"`

	// 本地镜像所关联服务器的IP（公网或私网）
	HostIp *string `json:"host_ip,omitempty"`

	// 本地镜像所关联容器的ID
	ContainerId *string `json:"container_id,omitempty"`

	// 本地镜像所关联容器的名称
	ContainerName *string `json:"container_name,omitempty"`

	// 本地镜像所关联Pod的ID
	PodId *string `json:"pod_id,omitempty"`

	// 本地镜像所关联Pod的名称
	PodName *string `json:"pod_name,omitempty"`

	// 本地镜像所关联软件的名称
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 是否存在容器 **约束限制**: 不涉及 **取值范围**: - true：是。 - false：否。  **默认取值**: 不涉及
	HasContainer *bool `json:"has_container,omitempty"`
}

func (o ListImageLocalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageLocalRequest struct{}"
	}

	return strings.Join([]string{"ListImageLocalRequest", string(data)}, " ")
}
