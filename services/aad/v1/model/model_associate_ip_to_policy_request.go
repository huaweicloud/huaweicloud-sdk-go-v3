package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIpToPolicyRequest Request Object
type AssociateIpToPolicyRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *IpBindingBody `json:"body,omitempty"`
}

func (o AssociateIpToPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIpToPolicyRequest struct{}"
	}

	return strings.Join([]string{"AssociateIpToPolicyRequest", string(data)}, " ")
}
