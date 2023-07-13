package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePeriodToOnDemandRequest Request Object
type UpdatePeriodToOnDemandRequest struct {
	Body *PeriodToOnDemandReq `json:"body,omitempty"`
}

func (o UpdatePeriodToOnDemandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeriodToOnDemandRequest struct{}"
	}

	return strings.Join([]string{"UpdatePeriodToOnDemandRequest", string(data)}, " ")
}
