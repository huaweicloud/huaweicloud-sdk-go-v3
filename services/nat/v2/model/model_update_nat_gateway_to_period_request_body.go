package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNatGatewayToPeriodRequestBody struct {
	PrepaidOptions *PrepaidOptions `json:"prepaid_options,omitempty"`
}

func (o UpdateNatGatewayToPeriodRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayToPeriodRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayToPeriodRequestBody", string(data)}, " ")
}
