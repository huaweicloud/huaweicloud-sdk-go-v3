package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryE2eSettingResponse Response Object
type ShowRepositoryE2eSettingResponse struct {
	E2ePolicies *E2ePolicyDto `json:"e2e_policies,omitempty"`

	Req *ReqSettingDto `json:"req,omitempty"`

	Link           *LinkSettingDto `json:"link,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowRepositoryE2eSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryE2eSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryE2eSettingResponse", string(data)}, " ")
}
