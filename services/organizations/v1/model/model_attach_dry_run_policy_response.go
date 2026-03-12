package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachDryRunPolicyResponse Response Object
type AttachDryRunPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachDryRunPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDryRunPolicyResponse struct{}"
	}

	return strings.Join([]string{"AttachDryRunPolicyResponse", string(data)}, " ")
}
