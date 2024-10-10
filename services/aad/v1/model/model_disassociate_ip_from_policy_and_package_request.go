package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateIpFromPolicyAndPackageRequest Request Object
type DisassociateIpFromPolicyAndPackageRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *IpBindingV3Body `json:"body,omitempty"`
}

func (o DisassociateIpFromPolicyAndPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateIpFromPolicyAndPackageRequest struct{}"
	}

	return strings.Join([]string{"DisassociateIpFromPolicyAndPackageRequest", string(data)}, " ")
}
