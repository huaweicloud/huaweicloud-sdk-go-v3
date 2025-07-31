package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksRequestBodyClusterScanInfo 集群扫描任务的参数，任务类型task_type为“cluster_scan”时可传
type ListTasksRequestBodyClusterScanInfo struct {

	// 任务扫描项类型列表，若列表不为空则只有扫描了列表中所有扫描项的任务才会被筛选出来；列表为空则不对扫描项类型进行筛选
	ScanTypeList *[]string `json:"scan_type_list,omitempty"`
}

func (o ListTasksRequestBodyClusterScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequestBodyClusterScanInfo struct{}"
	}

	return strings.Join([]string{"ListTasksRequestBodyClusterScanInfo", string(data)}, " ")
}
