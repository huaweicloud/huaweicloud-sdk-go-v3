package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeLoadbalancerChargeModeRequest struct {
	Body *ChangeLoadbalancerChargeModeRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ChangeLoadbalancerChargeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerChargeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerChargeModeRequest", string(data)}, " ")
}
