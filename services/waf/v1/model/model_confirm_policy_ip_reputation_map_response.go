package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmPolicyIpReputationMapResponse Response Object
type ConfirmPolicyIpReputationMapResponse struct {
	IpReputationMap *IpReputationMapResponseBodyIpReputationMap `json:"ip_reputation_map,omitempty"`

	Locale         *IpReputationMapResponseBodyLocale `json:"locale,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ConfirmPolicyIpReputationMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmPolicyIpReputationMapResponse struct{}"
	}

	return strings.Join([]string{"ConfirmPolicyIpReputationMapResponse", string(data)}, " ")
}
