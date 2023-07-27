package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportRequest Request Object
type ShowReportRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

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
