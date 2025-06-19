package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePeriodElasticResourcePoolSpecChangeOrderRequest Request Object
type CreatePeriodElasticResourcePoolSpecChangeOrderRequest struct {
	Body *CreatePeriodElasticResourcePoolSpecChangeOrderRequestBody `json:"body,omitempty"`
}

func (o CreatePeriodElasticResourcePoolSpecChangeOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePeriodElasticResourcePoolSpecChangeOrderRequest struct{}"
	}

	return strings.Join([]string{"CreatePeriodElasticResourcePoolSpecChangeOrderRequest", string(data)}, " ")
}
