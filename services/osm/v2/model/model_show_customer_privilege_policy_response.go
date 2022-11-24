package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCustomerPrivilegePolicyResponse struct {

	// 开关状态
	PolicySwitch   *bool `json:"policy_switch,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowCustomerPrivilegePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerPrivilegePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomerPrivilegePolicyResponse", string(data)}, " ")
}
