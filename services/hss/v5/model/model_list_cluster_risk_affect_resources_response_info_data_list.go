package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterRiskAffectResourcesResponseInfoDataList 集群风险影响的资源信息
type ListClusterRiskAffectResourcesResponseInfoDataList struct {

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源所属的命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 资源被检测出风险的命中规则
	HitRule *string `json:"hit_rule,omitempty"`

	// 资源存在风险的路径列表
	HitPathList *[]string `json:"hit_path_list,omitempty"`

	// 首次扫描时间
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// 最近扫描时间
	LastScanTime *int64 `json:"last_scan_time,omitempty"`
}

func (o ListClusterRiskAffectResourcesResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterRiskAffectResourcesResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"ListClusterRiskAffectResourcesResponseInfoDataList", string(data)}, " ")
}
