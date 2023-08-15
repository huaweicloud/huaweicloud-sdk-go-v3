package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstancesPeriodOrderRequest Request Object
type CreateInstancesPeriodOrderRequest struct {
	Body *CreateInstancePeriodRequest `json:"body,omitempty"`
}

func (o CreateInstancesPeriodOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancesPeriodOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateInstancesPeriodOrderRequest", string(data)}, " ")
}
