package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeLoadbalancerChargeModeRequest Request Object
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
