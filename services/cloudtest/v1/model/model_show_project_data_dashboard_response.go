package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectDataDashboardResponse Response Object
type ShowProjectDataDashboardResponse struct {
	Defect *DefectVo `json:"defect,omitempty"`

	CasePassRate *CasePassRateVo `json:"case_pass_rate,omitempty"`

	CaseCompletionRate *CaseCompletionRateVo `json:"case_completion_rate,omitempty"`

	// 项目类型
	ProjectType *string `json:"project_type,omitempty"`

	IssueCoverRate *IssueCoverRateVo `json:"issue_cover_rate,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowProjectDataDashboardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectDataDashboardResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectDataDashboardResponse", string(data)}, " ")
}
