package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskRequestBodyClusterScanInfo 创建集群扫描任务的参数，任务类型为“cluster_scan”时必传
type CreateTaskRequestBodyClusterScanInfo struct {

	// **参数解释**: 扫描项类型列表 **约束限制**: 不涉及 **取值范围**: - cluster_vul：集群漏洞 - risk_assessment：风险评估 - benchmark：安全合规  **默认取值**: 不涉及
	ScanTypeList []string `json:"scan_type_list"`

	// 扫描范围类型，包含如下   - all_cluster：扫描所有符合扫描条件的集群   - specific_cluster: 扫描指定集群
	RangeType string `json:"range_type"`

	// 需要扫描的集群id列表，扫描范围类型为“specific_cluster”时必传
	ClusterIdList *[]string `json:"cluster_id_list,omitempty"`
}

func (o CreateTaskRequestBodyClusterScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequestBodyClusterScanInfo struct{}"
	}

	return strings.Join([]string{"CreateTaskRequestBodyClusterScanInfo", string(data)}, " ")
}
