package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksResponseInfoClusterScanInfo 集群扫描任务特有的任务信息
type ListTasksResponseInfoClusterScanInfo struct {

	// 本次扫描任务涉及的扫描项类型列表
	ScanTypeList *[]string `json:"scan_type_list,omitempty"`

	// 当前任务处于扫描中的集群数量
	ScanningClusterNum *int32 `json:"scanning_cluster_num,omitempty"`

	// 当前任务扫描成功的集群数量
	SuccessClusterNum *int32 `json:"success_cluster_num,omitempty"`

	// 当前任务扫描失败的集群数量
	FailedClusterNum *int32 `json:"failed_cluster_num,omitempty"`
}

func (o ListTasksResponseInfoClusterScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponseInfoClusterScanInfo struct{}"
	}

	return strings.Join([]string{"ListTasksResponseInfoClusterScanInfo", string(data)}, " ")
}
