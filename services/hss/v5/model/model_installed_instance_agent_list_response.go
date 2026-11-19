package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstalledInstanceAgentListResponse Response Object
type InstalledInstanceAgentListResponse struct {

	// **参数解释**: 安装列表总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 安装列表接入异常工作负载/集群个数 **取值范围**: 最小值0，最大值2147483647
	AccessErrorNum *int32 `json:"access_error_num,omitempty"`

	// **参数解释**: 安装列表运行异常工作负载/集群个数 **取值范围**: 最小值0，最大值2147483647
	RunErrorNum *int32 `json:"run_error_num,omitempty"`

	// **参数解释**: 安装列表待升级Agent数 **取值范围**: 最小值0，最大值2147483647
	UpgradeNum *int32 `json:"upgrade_num,omitempty"`

	// **参数解释**: 安装列表升级失败实例数 **取值范围**: 最小值0，最大值2147483647
	UpgradeFailedNum *int32 `json:"upgrade_failed_num,omitempty"`

	// **参数解释**: 集群安装agent列表 **取值范围**: 最小值0，最大值65535
	ClusterList *[]ClusterInfo `json:"cluster_list,omitempty"`

	// **参数解释**: 工作负载安装agent列表 **取值范围**: 最小值0，最大值65535
	WorkloadList   *[]WorkLoadInfo `json:"workload_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o InstalledInstanceAgentListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstalledInstanceAgentListResponse struct{}"
	}

	return strings.Join([]string{"InstalledInstanceAgentListResponse", string(data)}, " ")
}
