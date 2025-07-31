package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResourcesRequestBody 查询当前任务关联的资源列表的请求体
type ListTaskResourcesRequestBody struct {

	// 任务类型，包含如下   - cluster_scan：集群扫描   - iac_scan：iac扫描
	TaskType string `json:"task_type"`

	ClusterScanInfo *ListTaskResourcesRequestBodyClusterScanInfo `json:"cluster_scan_info,omitempty"`

	IacScanInfo *ListTaskResourcesRequestBodyIacScanInfo `json:"iac_scan_info,omitempty"`
}

func (o ListTaskResourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResourcesRequestBody struct{}"
	}

	return strings.Join([]string{"ListTaskResourcesRequestBody", string(data)}, " ")
}
