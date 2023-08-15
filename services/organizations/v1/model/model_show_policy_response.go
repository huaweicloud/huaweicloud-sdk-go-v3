package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyResponse Response Object
type ShowPolicyResponse struct {
	Policy         *PolicyDto `json:"policy,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyResponse", string(data)}, " ")
}
