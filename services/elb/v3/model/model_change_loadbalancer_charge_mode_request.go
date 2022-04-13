package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeLoadbalancerChargeModeRequest struct {
	Body *ChangeLoadbalancerChargeModeRequestBody `json:"body,omitempty"`
}

func (o ChangeLoadbalancerChargeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerChargeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerChargeModeRequest", string(data)}, " ")
}
