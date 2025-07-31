package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskRequestBody 创建任务的请求体
type CreateTaskRequestBody struct {

	// 任务类型，包含如下   - cluster_scan：集群扫描   - iac_scan：iac扫描
	TaskType string `json:"task_type"`

	ClusterScanInfo *CreateTaskRequestBodyClusterScanInfo `json:"cluster_scan_info,omitempty"`

	IacScanInfo *CreateTaskRequestBodyIacScanInfo `json:"iac_scan_info,omitempty"`
}

func (o CreateTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTaskRequestBody", string(data)}, " ")
}
