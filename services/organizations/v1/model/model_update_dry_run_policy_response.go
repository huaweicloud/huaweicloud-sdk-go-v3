package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDryRunPolicyResponse Response Object
type UpdateDryRunPolicyResponse struct {
	Policy         *PolicyDto `json:"policy,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o UpdateDryRunPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDryRunPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDryRunPolicyResponse", string(data)}, " ")
}
