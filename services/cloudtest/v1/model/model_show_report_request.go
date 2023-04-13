package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowReportRequest struct {

	// 分支/计划id
	PlanId string `json:"plan_id"`

	Body *GenerateReportInfo `json:"body,omitempty"`
}

func (o ShowReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportRequest struct{}"
	}

	return strings.Join([]string{"ShowReportRequest", string(data)}, " ")
}
