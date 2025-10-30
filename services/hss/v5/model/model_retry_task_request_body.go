package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryTaskRequestBody 重新运行任务的请求体
type RetryTaskRequestBody struct {

	// **参数解释**： 任务类型 **约束限制**: 必填 **取值范围**: - cluster_scan：集群扫描。  **默认取值**: 不涉及
	TaskType string `json:"task_type"`

	ClusterScanInfo *RetryTaskRequestBodyClusterScanInfo `json:"cluster_scan_info,omitempty"`
}

func (o RetryTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryTaskRequestBody struct{}"
	}

	return strings.Join([]string{"RetryTaskRequestBody", string(data)}, " ")
}
