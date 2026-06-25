package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePeriodToOnDemandInstantlyRequest Request Object
type UpdatePeriodToOnDemandInstantlyRequest struct {
	Body *PeriodToOnDemandInstantlyReq `json:"body,omitempty"`
}

func (o UpdatePeriodToOnDemandInstantlyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeriodToOnDemandInstantlyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePeriodToOnDemandInstantlyRequest", string(data)}, " ")
}
