package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePeriodToOnDemandRequest struct {
	Body *PeriodToOnDemandReq `json:"body,omitempty" xml:"body"`
}

func (o UpdatePeriodToOnDemandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeriodToOnDemandRequest struct{}"
	}

	return strings.Join([]string{"UpdatePeriodToOnDemandRequest", string(data)}, " ")
}
