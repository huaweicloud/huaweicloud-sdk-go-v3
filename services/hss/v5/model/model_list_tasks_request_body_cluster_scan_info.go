package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksRequestBodyClusterScanInfo 集群扫描任务的参数，任务类型task_type为“cluster_scan”时可传
type ListTasksRequestBodyClusterScanInfo struct {

	// **参数解释**: 任务扫描项类型列表，若列表不为空则只有扫描了列表中所有扫描项的任务才会被筛选出来；列表为空则不对扫描项类型进行筛选 **约束限制**: 不涉及 **取值范围**: - cluster_vul：集群漏洞 - risk_assessment：风险评估 - benchmark：安全合规  **默认取值**: 不涉及
	ScanTypeList *[]string `json:"scan_type_list,omitempty"`
}

func (o ListTasksRequestBodyClusterScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequestBodyClusterScanInfo struct{}"
	}

	return strings.Join([]string{"ListTasksRequestBodyClusterScanInfo", string(data)}, " ")
}
