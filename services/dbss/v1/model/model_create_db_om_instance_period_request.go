package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbOmInstancePeriodRequest Request Object
type CreateDbOmInstancePeriodRequest struct {
	Body *ChargeOrderDbssNew `json:"body,omitempty"`
}

func (o CreateDbOmInstancePeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbOmInstancePeriodRequest struct{}"
	}

	return strings.Join([]string{"CreateDbOmInstancePeriodRequest", string(data)}, " ")
}
