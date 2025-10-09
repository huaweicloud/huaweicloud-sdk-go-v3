package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetClusterQuotaBodyQuotas **参数解释**： 集群配额 **约束限制**： 不涉及
type GetClusterQuotaBodyQuotas struct {

	// **参数解释**： 集群配额 **约束限制**： 不涉及
	Resources *[]ClusterQuotaResource `json:"resources,omitempty"`
}

func (o GetClusterQuotaBodyQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClusterQuotaBodyQuotas struct{}"
	}

	return strings.Join([]string{"GetClusterQuotaBodyQuotas", string(data)}, " ")
}
