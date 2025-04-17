package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDefensePolicyRequest Request Object
type EnableDefensePolicyRequest struct {

	// 用户EIP对应的ID
	FloatingIpId string `json:"floating_ip_id"`

	Body *OpenAntiDDosServiceRequestBody `json:"body,omitempty"`
}

func (o EnableDefensePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDefensePolicyRequest struct{}"
	}

	return strings.Join([]string{"EnableDefensePolicyRequest", string(data)}, " ")
}
