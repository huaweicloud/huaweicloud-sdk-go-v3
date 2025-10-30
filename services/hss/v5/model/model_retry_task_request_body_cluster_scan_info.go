package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryTaskRequestBodyClusterScanInfo task_type值为“cluster_scan”时用于传递重新执行集群扫描任务的相关参数
type RetryTaskRequestBodyClusterScanInfo struct {

	// **参数解释**： 重新扫描的范围 **约束限制**: 必填 **取值范围**: - all_failed_cluster：任务下所有扫描失败的集群。 - specific_cluster：任务下指定扫描失败的集群。  **默认取值**: 不涉及
	RangeType string `json:"range_type"`

	// **参数解释**： 扫描范围为specific_cluster时用于指定重新扫描的集群范围 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值200
	ClusterInfoList *[]RetryTaskRequestBodyClusterScanInfoClusterInfoList `json:"cluster_info_list,omitempty"`
}

func (o RetryTaskRequestBodyClusterScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryTaskRequestBodyClusterScanInfo struct{}"
	}

	return strings.Join([]string{"RetryTaskRequestBodyClusterScanInfo", string(data)}, " ")
}
