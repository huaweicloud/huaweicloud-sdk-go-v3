package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectE2eSettingResponse Response Object
type ShowProjectE2eSettingResponse struct {
	E2ePolicies *E2ePolicyDto `json:"e2e_policies,omitempty"`

	Req *ReqSettingDto `json:"req,omitempty"`

	Link           *LinkSettingDto `json:"link,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowProjectE2eSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectE2eSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectE2eSettingResponse", string(data)}, " ")
}
