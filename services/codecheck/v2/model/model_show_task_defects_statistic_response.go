package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskDefectsStatisticResponse struct {
	Severity *StatisticSeverityV2 `json:"severity,omitempty"`

	Status         *StatisticStatusV2 `json:"status,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowTaskDefectsStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDefectsStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDefectsStatisticResponse", string(data)}, " ")
}
