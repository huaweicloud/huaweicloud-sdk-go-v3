package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIpToPolicyAndPackageRequest Request Object
type AssociateIpToPolicyAndPackageRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *IpBindingV3Body `json:"body,omitempty"`
}

func (o AssociateIpToPolicyAndPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIpToPolicyAndPackageRequest struct{}"
	}

	return strings.Join([]string{"AssociateIpToPolicyAndPackageRequest", string(data)}, " ")
}
