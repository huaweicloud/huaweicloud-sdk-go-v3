package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskTypeSummaryResponse Response Object
type ShowTaskTypeSummaryResponse struct {
	Filter *ShowTaskTypeSummaryResponseBodyFilter `json:"filter,omitempty"`

	Summary *ShowTaskTypeSummaryResponseBodySummary `json:"summary,omitempty"`

	// 打点数据的条数
	Count *int64 `json:"count,omitempty"`

	// 打点数据
	Datapoints     *[]ShowTaskTypeSummaryResponseBodyDatapoints `json:"datapoints,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ShowTaskTypeSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskTypeSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskTypeSummaryResponse", string(data)}, " ")
}
