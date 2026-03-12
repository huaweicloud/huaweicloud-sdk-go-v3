package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDryRunPolicyResponse Response Object
type ShowDryRunPolicyResponse struct {
	Policy         *PolicyDto `json:"policy,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowDryRunPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDryRunPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDryRunPolicyResponse", string(data)}, " ")
}
