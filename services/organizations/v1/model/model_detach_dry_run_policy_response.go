package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachDryRunPolicyResponse Response Object
type DetachDryRunPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachDryRunPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachDryRunPolicyResponse struct{}"
	}

	return strings.Join([]string{"DetachDryRunPolicyResponse", string(data)}, " ")
}
