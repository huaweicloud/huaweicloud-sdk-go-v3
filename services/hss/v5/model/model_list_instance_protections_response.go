package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceProtectionsResponse Response Object
type ListInstanceProtectionsResponse struct {

	// **参数解释**: 防护列表总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 集群防护列表 **取值范围**: 最小值0，最大值65535
	ClusterList *[]ClusterInfo `json:"cluster_list,omitempty"`

	// **参数解释**: 工作负载防护列表 **取值范围**: 最小值0，最大值65535
	WorkloadList *[]WorkLoadInfo `json:"workload_list,omitempty"`

	// **参数解释**: 实例防护列表 **取值范围**: 最小值0，最大值65535
	InstanceList *[]ServerlessInstanceInfo `json:"instance_list,omitempty"`

	// **参数解释**: 有风险实例数 **取值范围**: 最小值0，最大值2147483647
	RiskNum *int32 `json:"risk_num,omitempty"`

	// **参数解释**: 防护异常数 **取值范围**: 最小值0，最大值2147483647
	AbnormalNum *int32 `json:"abnormal_num,omitempty"`

	// **参数解释**: 未防护实例数 **取值范围**: 最小值0，最大值2147483647
	UnprotectNum *int32 `json:"unprotect_num,omitempty"`

	// **参数解释**: 全部实例数 **取值范围**: 最小值0，最大值2147483647
	InstanceNum    *int32 `json:"instance_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceProtectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceProtectionsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceProtectionsResponse", string(data)}, " ")
}
