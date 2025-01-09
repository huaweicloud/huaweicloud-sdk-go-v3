package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyPolicyRequest struct {
	Policies *Policies `json:"policies,omitempty"`
}

func (o ModifyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPolicyRequest struct{}"
	}

	return strings.Join([]string{"ModifyPolicyRequest", string(data)}, " ")
}
