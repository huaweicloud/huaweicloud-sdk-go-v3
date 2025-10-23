package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskStatusSummaryResponse Response Object
type ShowTaskStatusSummaryResponse struct {
	Filter *ShowTaskStatusSummaryResponseBodyFilter `json:"filter,omitempty"`

	Summary *ShowTaskStatusSummaryResponseBodySummary `json:"summary,omitempty"`

	// 打点数据的条数
	Count *int64 `json:"count,omitempty"`

	// 打点数据
	Datapoints     *[]ShowTaskStatusSummaryResponseBodyDatapoints `json:"datapoints,omitempty"`
	HttpStatusCode int                                            `json:"-"`
}

func (o ShowTaskStatusSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskStatusSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskStatusSummaryResponse", string(data)}, " ")
}
