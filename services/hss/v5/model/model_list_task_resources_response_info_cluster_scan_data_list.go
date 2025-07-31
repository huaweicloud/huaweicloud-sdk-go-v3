package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResourcesResponseInfoClusterScanDataList 任务下的集群扫描详情信息
type ListTaskResourcesResponseInfoClusterScanDataList struct {

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群类型，包含如下： - cce：CCE集群 - ali：阿里云集群 - tencent：腾讯云集群 - azure：微软云集群 - aws：亚马逊集群 - self_built_hw：华为云自建集群 - self_built_idc：IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`

	// 集群版本
	ClusterVersion *string `json:"cluster_version,omitempty"`

	// 集群的扫描状态，包含如下：   - scanning：扫描中   - success：扫描成功   - failed：扫描失败
	ScanStatus *string `json:"scan_status,omitempty"`

	// 该集群扫描开始的时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 该集群扫描结束的时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 当前集群的扫描详情信息
	ScanDetailList *[]ListTaskResourcesResponseInfoScanDetailList `json:"scan_detail_list,omitempty"`
}

func (o ListTaskResourcesResponseInfoClusterScanDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResourcesResponseInfoClusterScanDataList struct{}"
	}

	return strings.Join([]string{"ListTaskResourcesResponseInfoClusterScanDataList", string(data)}, " ")
}
