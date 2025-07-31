package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResourcesResponse Response Object
type ListTaskResourcesResponse struct {

	// 任务关联的资源数量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 本次任务关联的集群扫描详情列表，任务类型为“cluster_scan”时查询结果为该列表
	ClusterScanDataList *[]ListTaskResourcesResponseInfoClusterScanDataList `json:"cluster_scan_data_list,omitempty"`

	// 本次任务关联的文件扫描详情列表，任务类型为“iac_scan”时查询结果为该列表
	IacScanDataList *[]ListTaskResourcesResponseInfoIacScanDataList `json:"iac_scan_data_list,omitempty"`
	HttpStatusCode  int                                             `json:"-"`
}

func (o ListTaskResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListTaskResourcesResponse", string(data)}, " ")
}
