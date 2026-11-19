package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityRiskResponse Response Object
type ShowSecurityRiskResponse struct {
	HostRiskInfo *HostRiskInfo `json:"host_risk_info,omitempty"`

	ContainerRiskInfo *ContainerRiskInfo `json:"container_risk_info,omitempty"`

	// **参数解释**: 安全风险趋势 **取值范围**: 取值0-30个TrendRiskInfo对象
	TrendDataList  *[]TrendRiskInfo `json:"trend_data_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowSecurityRiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityRiskResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityRiskResponse", string(data)}, " ")
}
